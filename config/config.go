package config

type Config struct {
	Port      string
	Addr      string
	SecretKey string
	Db        DB
}

type DB struct {
	Port     string
	Host     string
	Database string
	User     string
	Password string
}

// rpc RaiseDisputeOnOrder(RaiseDisputeOnOrderRequest) returns (RaiseDisputeOnOrderResponse){} Future service
