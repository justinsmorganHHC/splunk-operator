package common

// PATHS
// List of all Paths used in the Splunk Operator

const (

	//*****************
	// Splunk Enterprise Paths
	//*****************

	//PeerApps path
	PeerApps = "etc/slave-apps"

	//ManagerApps = "etc/master-apps"
	ManagerApps = "etc/master-apps"

	//SHCApps = "etc/shcluster/apps"
	SHCApps = "etc/shcluster/apps"

	//*****************
	// Kubernetes Paths
	//*****************

	//ClusterManagerService = "cluster-master-service"
	ClusterManagerService = CM + Dash + Service

	//ClusterManagerSecret = "cluster-master-secret"
	ClusterManagerSecret = CM + Dash + Secret

	//SplunkExample = "splunk-example"
	SplunkExample = Spl + Dash + "example"

	//TestSplunkStack1 = "-test-splunk-stack1"
	TestSplunkStack1 = Dash + Test + Dash + Spl + Dash + Stack1

	//ServiceTestStack1 = "Service-test-splunk-stack1"
	ServiceTestStack1 = CapitalService + TestSplunkStack1

	//SecretTestStack1 = "Secret-test-splunk-stack1"
	SecretTestStack1 = CapitalSecret + TestSplunkStack1

	//ConfigMapTestStack1 = "ConfigMap-test-splunk-stack1"
	ConfigMapTestStack1 = ConfigM + TestSplunkStack1

	//StatefulSetStack1 = "StatefulSet-test-splunk-stack1"
	StatefulSetStack1 = StatefulS + TestSplunkStack1

	//StatefulSetCMTest = "StatefulSet-test-splunk-stack1-cluster-master"
	StatefulSetCMTest = StatefulSetStack1 + Dash + CM

	//ConfigMCMTestSmartStore = "ConfigMap-test-splunk-stack1-clustermaster-smartstore"
	ConfigMCMTestSmartStore = ConfigMapTestStack1 + Dash + ClusterM + Dash + SmartStore

	//ConfigMCMTestAppList = "ConfigMap-test-splunk-stack1-clustermaster-app-list"
	ConfigMCMTestAppList = ConfigMapTestStack1 + Dash + ClusterM + Dash + AppList

	//CMStack1Service = "Service-test-splunk-stack1-cluster-master-service"
	CMStack1Service = ServiceTestStack1 + Dash + ClusterManagerService

	//CMStack1Secret = "Secret-test-splunk-stack1-cluster-master-secret-v1"
	CMStack1Secret = SecretTestStack1 + Dash + ClusterManagerSecret + Dash + V1

	//ClusterManagerExample = "splunk-example-cluster-master-service:8089"
	ClusterManagerExample = SplunkExample + Dash + ClusterManagerService + Port8089
)

// Base K8s Paths
const (

	//LM = "license-master"
	LM = "license-master"

	//CM = "cluster-master"
	CM = "cluster-master"

	//ClusterM = "clustermaster"  (Exceptional case)
	ClusterM = "clustermaster"

	//ConfigM = "ConfigMap"
	ConfigM = "ConfigMap"

	//StatefulS = "StatefulSet"
	StatefulS = "StatefulSet"

	//Secret = "secret"
	Secret = "secret"

	//CapitalSecret = "Secret"
	CapitalSecret = "Secret"

	//Service = "service"
	Service = "service"

	//CapitalService = "Service"
	CapitalService = "Service"

	//Test = "test"
	Test = "test"

	//Stack1 = "stack1"
	Stack1 = "stack1"

	//SmartStore = "smartstore"
	SmartStore = "smartstore"

	//AppList = "app-list"
	AppList = "app-list"

	//V1 = "v1"
	V1 = "v1"
)
