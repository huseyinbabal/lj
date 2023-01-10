package app

type Java struct {
	version string
}

func (app *App) GetJava() (Java, error) {
	version, err := app.jvm.Version()
	if err != nil {
		return Java{}, err
	}
	return Java{
		version: version,
	}, nil
}
