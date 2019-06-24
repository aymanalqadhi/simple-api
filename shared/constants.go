package shared

const (
	// ListenPort is the tcp port on which to start the server
	ListenPort = 3434

	// DBDriver is the databse server driver name
	DBDriver = "mysql"

	// DBaseHost is the host address of the database server
	// TODO: Sotre in environment variables for security
	DBaseHost = "localhost"

	// DBPort is the TCP port on which to connect
	DBPort = 3306

	// DBUsername is the username of the database server
	// TODO: Sotre in environment variables for security
	DBUsername = "root"

	// DBPassword is the password of the database server
	// TODO: Sotre in environment variables for security
	DBPassword = "123"

	// DBName is the name of the database
	// TODO: Sotre in environment variables for security
	DBName = "test"

	// DBOptions is the options string of the connection
	// TODO: Sotre in environment variables for security
	DBOptions = "parseTime=true"

	// AuthPassword is the password which used in authenticating users
	// TODO: Sotre in environment variables for security
	AuthPassword = "password123"
)
