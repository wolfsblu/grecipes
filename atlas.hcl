// Define an environment named "local"
env "local" {
  // Declare where the schema definition resides.
  // Also supported: ["file://multi.hcl", "file://schema.hcl"].
  src = "file://infra/sqlite/schema.sql"

  // Define the URL of the database which is managed
  // in this environment.
  url = "sqlite://tmp/db.sqlite"

  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = "sqlite://dev?mode=memory"

  migration {
    dir = "file://infra/sqlite/migrations"
  }
}