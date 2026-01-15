#!/bin/bash
set -e  # Faz o script parar se qualquer comando falhar

# Variáveis
DB_CONTAINER="url_shortener_db"
DB_USER="postgres"
DB_NAME="shortener"
RESET_SQL="./db/reset.sql"

echo "🚀 Limpando o banco de dados..."

# Executa o reset dentro do container do Postgres
docker exec -i $DB_CONTAINER psql -U $DB_USER -d $DB_NAME < $RESET_SQL

if [ $? -eq 0 ]; then
    echo "✅ Banco limpo com sucesso!"
    echo "🏃‍♂️ Iniciando teste de performance com K6..."

    # Rodando K6 na mesma network do docker-compose
    k6 run - < tests/performance/k6/shorten.js

    echo "🎉 Teste finalizado!"
else
    echo "❌ Falha ao limpar o banco. Abortando teste."
    exit 1
fi
