data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "ariga.io/atlas-provider-gorm",
    "load",
    "--path", "./model",
    "--dialect", "sqlite",
  ]
}

env "local" {
  src = data.external_schema.gorm.url
  dev = "sqlite://dev.db?mode=memory"
  url = "sqlite://mydb.sqlite"
  
  migration {
    dir = "file://migrations"
  }
}