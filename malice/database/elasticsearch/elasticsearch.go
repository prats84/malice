package elasticsearch

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/docker/docker/api/types"
	"github.com/docker/go-connections/nat"
	"github.com/maliceio/go-plugin-utils/utils"
	"github.com/maliceio/go-plugin-utils/waitforit"
	"github.com/maliceio/malice/config"
	"github.com/maliceio/malice/malice/database"
	"github.com/maliceio/malice/malice/docker/client"
	"github.com/maliceio/malice/malice/docker/client/container"
	er "github.com/maliceio/malice/malice/errors"
	elastic "gopkg.in/olivere/elastic.v5"
)

// PluginResults a malice plugin results object
type PluginResults struct {
	ID       string `json:"id"`
	Name     string
	Category string
	Data     map[string]interface{}
}

// ElasticAddr ElasticSearch address to user for connections
var ElasticAddr string

func getElasticSearchAddr(addr string) string {
	if _, inDocker := os.LookupEnv("MALICE_IN_DOCKER"); inDocker {
		if addr != "" {
			return fmt.Sprintf("http://%s:9200", utils.Getopt("MALICE_ELASTICSEARCH", addr))
		}
		return fmt.Sprintf("http://%s:9200", utils.Getopt("MALICE_ELASTICSEARCH", "elastic"))
	}
	return fmt.Sprintf("http://%s:9200", utils.Getopt("MALICE_ELASTICSEARCH", "localhost"))
}

// StartELK creates an ELK container from the image blacktop/elk
func StartELK(docker *client.Docker, logs bool) (types.ContainerJSONBase, error) {

	name := config.Conf.DB.Name
	image := config.Conf.DB.Image
	binds := []string{"malice:/usr/share/elasticsearch/data"}
	portBindings := nat.PortMap{
		"80/tcp":   {{HostIP: "0.0.0.0", HostPort: "80"}},
		"9200/tcp": {{HostIP: "0.0.0.0", HostPort: "9200"}},
	}

	if docker.Ping() {
		cont, err := container.Start(docker, nil, name, image, logs, binds, portBindings, nil, nil)

		// Inspect newly created container to get IP assigned to it
		dbInfo, err := container.Inspect(docker, cont.ID)
		elasticAddress := getElasticSearchAddr(dbInfo.NetworkSettings.IPAddress)

		// Give ELK a few seconds to start
		log.WithFields(log.Fields{
			"server":  elasticAddress,
			"timeout": config.Conf.DB.Timeout,
		}).Debug("Waiting for ELK to come online.")
		if err = waitforit.WaitForIt(elasticAddress, "", -1, config.Conf.DB.Timeout); err != nil {
			log.Error(err)
		}
		log.Debug("ELK is now online.")

		// Even though it's up it's not ready to index data yet.
		log.Infof("Sleeping for 10 seconds to give %s time to initalize.", config.Conf.DB.Image)
		time.Sleep(10 * time.Second)

		return cont, err
	}
	return types.ContainerJSONBase{}, errors.New("Cannot connect to the Docker daemon. Is the docker daemon running on this host?")
}

// InitElasticSearch initalizes ElasticSearch for use with malice
func InitElasticSearch(addr string) error {

	// Test connection to ElasticSearch
	er.CheckError(TestConnection(addr))

	client, err := elastic.NewSimpleClient(elastic.SetURL(ElasticAddr))
	utils.Assert(err)

	exists, err := client.IndexExists("malice").Do(context.Background())
	utils.Assert(err)

	if !exists {
		// Index does not exist yet.
		createIndex, err := client.CreateIndex("malice").BodyString(mapping).Do(context.Background())
		utils.Assert(err)
		if !createIndex.Acknowledged {
			// Not acknowledged
			log.Error("Couldn't create Index.")
		} else {
			log.Debug("Created Index: ", "malice")
		}
	} else {
		log.Debug("Index malice already exists.")
	}

	return err
}

// TestConnection tests the ElasticSearch connection
func TestConnection(addr string) error {

	var err error

	if ElasticAddr == "" {
		ElasticAddr = getElasticSearchAddr(addr)
	}

	// connect to ElasticSearch where --link elastic was using via malice in Docker
	log.Debugf("Attempting to connect to: %s", ElasticAddr)
	client, err := elastic.NewSimpleClient(elastic.SetURL(ElasticAddr))

	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping(ElasticAddr).Do(context.Background())
	utils.Assert(err)

	log.WithFields(log.Fields{
		"code":    code,
		"cluster": info.ClusterName,
		"version": info.Version.Number,
		"address": ElasticAddr,
	}).Debug("ElasticSearch connection successful.")

	return err
}

// WriteFileToDatabase inserts sample into Database
func WriteFileToDatabase(sample map[string]interface{}) elastic.IndexResponse {

	// Test connection to ElasticSearch
	er.CheckError(TestConnection(""))

	client, err := elastic.NewSimpleClient(elastic.SetURL(ElasticAddr))
	utils.Assert(err)

	scan := map[string]interface{}{
		// "id":      sample.SHA256,
		"file":      sample,
		"plugins":   database.GetPluginsByCategory(),
		"scan_date": time.Now().Format(time.RFC3339Nano),
	}

	newScan, err := client.Index().
		Index("malice").
		Type("samples").
		OpType("index").
		// Id("1").
		BodyJson(scan).
		Do(context.Background())
	utils.Assert(err)

	log.WithFields(log.Fields{
		"id":    newScan.Id,
		"index": newScan.Index,
		"type":  newScan.Type,
	}).Debug("Indexed sample.")

	return *newScan
}

// WriteHashToDatabase inserts sample into Database
func WriteHashToDatabase(hash string) elastic.IndexResponse {

	hashType, err := utils.GetHashType(hash)
	utils.Assert(err)

	client, err := elastic.NewSimpleClient(elastic.SetURL(ElasticAddr))
	utils.Assert(err)

	scan := map[string]interface{}{
		// "id":      sample.SHA256,
		"file": map[string]interface{}{
			hashType: hash,
		},
		"plugins":   database.GetPluginsByCategory(),
		"scan_date": time.Now().Format(time.RFC3339Nano),
	}

	newScan, err := client.Index().
		Index("malice").
		Type("samples").
		OpType("create").
		// Id("1").
		BodyJson(scan).
		Do(context.Background())
	utils.Assert(err)

	log.WithFields(log.Fields{
		"id":    newScan.Id,
		"index": newScan.Index,
		"type":  newScan.Type,
	}).Debug("Indexed sample.")

	return *newScan
}

// WritePluginResultsToDatabase upserts plugin results into Database
func WritePluginResultsToDatabase(results PluginResults) {

	// scanID := utils.Getopt("MALICE_SCANID", "")
	if ElasticAddr == "" {
		ElasticAddr = fmt.Sprintf("http://%s:9200", utils.Getopt("MALICE_ELASTICSEARCH", "elastic"))
	}

	client, err := elastic.NewSimpleClient(elastic.SetURL(ElasticAddr))
	utils.Assert(err)

	getSample, err := client.Get().
		Index("malice").
		Type("samples").
		Id(results.ID).
		Do(context.Background())

	if err != nil {

	}

	if getSample != nil && getSample.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", getSample.Id, getSample.Version, getSample.Index, getSample.Type)
		updateScan := map[string]interface{}{
			"scan_date": time.Now().Format(time.RFC3339Nano),
			"plugins": map[string]interface{}{
				results.Category: map[string]interface{}{
					results.Name: results.Data,
				},
			},
		}
		update, err := client.Update().Index("malice").Type("samples").Id(getSample.Id).
			Doc(updateScan).
			Do(context.Background())
		utils.Assert(err)

		log.WithFields(log.Fields{
			"id":      update.Id,
			"version": update.Version,
		}).Debug("New version of sample.")

		// return *update

	} else {

		scan := map[string]interface{}{
			// "id":      sample.SHA256,
			// "file":      sample,
			"plugins": map[string]interface{}{
				results.Category: map[string]interface{}{
					results.Name: results.Data,
				},
			},
			"scan_date": time.Now().Format(time.RFC3339Nano),
		}

		newScan, err := client.Index().
			Index("malice").
			Type("samples").
			OpType("create").
			// Id("1").
			BodyJson(scan).
			Do(context.Background())
		utils.Assert(err)

		log.WithFields(log.Fields{
			"id":    newScan.Id,
			"index": newScan.Index,
			"type":  newScan.Type,
		}).Debug("Indexed sample.")
		// return *newScan
	}
}
