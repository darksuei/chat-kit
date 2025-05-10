if [ -z "$1" ]; then
  echo "Error (Missing required argument): migration_name."
  echo "Usage: $0 migration_name"
  exit 1
fi

migrate create -ext sql -dir migrations -seq $1