package linodego_test

import (
	"fmt"
	"log"

	"github.com/chiefy/linodego"
)

func ExampleClient_CreateNodeBalancer() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleCreateNodeBalancer")
	defer teardown()

	fmt.Println("## NodeBalancer create")
	var nbID int
	var nb = &linodego.NodeBalancer{
		ClientConnThrottle: 20,
		Region:             "us-east",
	}

	createOpts := nb.GetCreateOptions()
	nb, err := linodeClient.CreateNodeBalancer(&createOpts)
	if err != nil {
		log.Fatal(err)
	}
	nbID = nb.ID

	fmt.Println("### Get")
	nb, err = linodeClient.GetNodeBalancer(nbID)
	if err != nil {
		log.Fatal(err)
	}

	updateOpts := nb.GetUpdateOptions()
	*updateOpts.Label += "_renamed"
	nb, err = linodeClient.UpdateNodeBalancer(nbID, updateOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Delete")
	if err := linodeClient.DeleteNodeBalancer(nbID); err != nil {
		log.Fatal(err)
	}

	// Output:
	// ## NodeBalancer create
	// ### Get
	// ### Delete
}

func ExampleClient_CreateNodeBalancerConfig() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleCreateNodeBalancerConfig")
	defer teardown()

	fmt.Println("## NodeBalancer create")
	clientConnThrottle := 20
	nb, err := linodeClient.CreateNodeBalancer(&linodego.NodeBalancerCreateOptions{
		ClientConnThrottle: &clientConnThrottle,
		Region:             "us-east",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("## NodeBalancer Config create")
	createOpts := linodego.NodeBalancerConfigCreateOptions{
		Port: 80,
		/*
			Protocol:      linodego.ProtocolHTTP,
			Algorithm:     linodego.AlgorithmLeastConn,
			Stickiness:    linodego.StickinessHTTPCookie,
			Check:         linodego.CheckHTTP,
			CheckInterval: 30,
			CheckAttempts: 5,
			CipherSuite:   linodego.CipherRecommended,
		*/
	}
	nbc, err := linodeClient.CreateNodeBalancerConfig(nb.ID, &createOpts)
	if err != nil {
		log.Fatal(err)
	}
	nbcID := nbc.ID

	fmt.Println("## NodeBalancer Config update")
	updateOpts := nbc.GetUpdateOptions()
	updateOpts.Port += 8000
	nbc, err = linodeClient.UpdateNodeBalancerConfig(nb.ID, nbc.ID, updateOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### List")
	configs, err := linodeClient.ListNodeBalancerConfigs(nb.ID, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Get")
	nbc, err = linodeClient.GetNodeBalancerConfig(nb.ID, configs[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Delete")
	if nbc.ID != nbcID {
		log.Fatalf("Unexpected Nodebalancer Config ID %d != %d", nbc.ID, nbcID)
	}
	if err := linodeClient.DeleteNodeBalancerConfig(nb.ID, nbc.ID); err != nil {
		log.Fatal(err)
	}

	if err := linodeClient.DeleteNodeBalancer(nb.ID); err != nil {
		log.Fatal(err)
	}

	// Output:
	// ## NodeBalancer create
	// ## NodeBalancer Config create
	// ## NodeBalancer Config update
	// ### List
	// ### Get
	// ### Delete
}

func ExampleClient_CreateNodeBalancerNode() {
	// Example readers, Ignore this bit of setup code needed to record test fixtures
	linodeClient, teardown := createTestClient(nil, "fixtures/ExampleCreateNodeBalancerNode")
	defer teardown()

	fmt.Println("## NodeBalancer create")
	clientConnThrottle := 20
	nb, err := linodeClient.CreateNodeBalancer(&linodego.NodeBalancerCreateOptions{
		ClientConnThrottle: &clientConnThrottle,
		Region:             "us-east",
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("## NodeBalancer Config create")

	nbc, err := linodeClient.CreateNodeBalancerConfig(nb.ID, &linodego.NodeBalancerConfigCreateOptions{
		Port: 80,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("## NodeBalancer Node create")
	createOpts := linodego.NodeBalancerNodeCreateOptions{
		Address: "192.168.129.255:80",
		Label:   "192.168.129.255-80",
	}
	nbn, err := linodeClient.CreateNodeBalancerNode(nb.ID, nbc.ID, &createOpts)
	if err != nil {
		log.Fatal(err)
	}
	nbnID := nbn.ID

	fmt.Println("## NodeBalancer Node update")
	updateOpts := nbn.GetUpdateOptions()
	updateOpts.Address = "192.168.129.0:8080"
	nbn, err = linodeClient.UpdateNodeBalancerNode(nb.ID, nbc.ID, nbn.ID, updateOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### List")
	nodes, err := linodeClient.ListNodeBalancerNodes(nb.ID, nbc.ID, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Get")
	nbn, err = linodeClient.GetNodeBalancerNode(nb.ID, nbc.ID, nodes[0].ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("### Delete")
	if nbn.ID != nbnID {
		log.Fatalf("Unexpected Nodebalancer Node ID %d != %d", nbn.ID, nbnID)
	}
	if err := linodeClient.DeleteNodeBalancerNode(nb.ID, nbc.ID, nbn.ID); err != nil {
		log.Fatal(err)
	}

	if err := linodeClient.DeleteNodeBalancerConfig(nb.ID, nbc.ID); err != nil {
		log.Fatal(err)
	}

	if err := linodeClient.DeleteNodeBalancer(nb.ID); err != nil {
		log.Fatal(err)
	}

	// Output:
	// ## NodeBalancer create
	// ## NodeBalancer Config create
	// ## NodeBalancer Node create
	// ## NodeBalancer Node update
	// ### List
	// ### Get
	// ### Delete
}
