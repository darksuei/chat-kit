if [ -z "$1" ]; then
  echo "Error (missing required argument): database_connection_str."
  echo "Usage: $0 database_connection_str"
  exit 1
fi

migrate -path internal/migrations -database $1 up