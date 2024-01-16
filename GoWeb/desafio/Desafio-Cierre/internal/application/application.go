package application

// Application is an interface that represent a application
type Application interface {
	// Setup is a method that setup the application
	Setup() (err error)
	// Run is a method that run the application
	Run() (err error)
}
