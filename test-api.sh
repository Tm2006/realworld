#!/bin/bash

# 🧪 RealWorld API Test Script
# Автоматизированное тестирование API endpoints

# Цвета для вывода
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

API_URL="http://localhost:8080"

echo -e "${BLUE}🧪 Testing RealWorld API${NC}"
echo "================================"

# Функция для тестирования endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local expected_status=$3
    local description=$4
    local data=$5
    
    echo -n "Testing $method $endpoint... "
    
    if [ -n "$data" ]; then
        response=$(curl -s -w "%{http_code}" -X $method "$API_URL$endpoint" \
                   -H "Content-Type: application/json" \
                   -d "$data")
    else
        response=$(curl -s -w "%{http_code}" -X $method "$API_URL$endpoint")
    fi
    
    status_code="${response: -3}"
    body="${response%???}"
    
    if [ "$status_code" -eq "$expected_status" ]; then
        echo -e "${GREEN}✅ PASS${NC} ($status_code)"
        if [ -n "$body" ] && [ ${#body} -gt 0 ]; then
            echo -e "   Response: ${body:0:100}..."
        fi
    else
        echo -e "${RED}❌ FAIL${NC} (expected $expected_status, got $status_code)"
        if [ -n "$body" ] && [ ${#body} -gt 0 ]; then
            echo -e "   Response: ${body:0:100}..."
        fi
    fi
}

# Проверка доступности сервера
echo -e "\n${YELLOW}🔍 Server Connectivity${NC}"
if curl -s --connect-timeout 5 "$API_URL" > /dev/null; then
    echo -e "Server connectivity: ${GREEN}✅ PASS${NC}"
else
    echo -e "Server connectivity: ${RED}❌ FAIL${NC}"
    echo "❗ Make sure the server is running: go run cmd/realworld/main.go"
    exit 1
fi

# CORS тесты
echo -e "\n${YELLOW}🌐 CORS Headers${NC}"
cors_response=$(curl -s -I -X OPTIONS "$API_URL/api/user" \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: GET")

if echo "$cors_response" | grep -q "Access-Control-Allow-Origin"; then
    echo -e "CORS headers: ${GREEN}✅ PASS${NC}"
    echo "   Origin: $(echo "$cors_response" | grep "Access-Control-Allow-Origin" | tr -d '\r')"
else
    echo -e "CORS headers: ${RED}❌ FAIL${NC}"
fi

# User endpoints
echo -e "\n${YELLOW}👤 User Endpoints${NC}"
test_endpoint "GET" "/api/user" 200 "Get current user"
test_endpoint "PUT" "/api/user" 200 "Update user" \
    '{"user": {"email": "test@example.com", "username": "testuser"}}'

# Articles endpoints
echo -e "\n${YELLOW}📝 Articles Endpoints${NC}"
test_endpoint "GET" "/api/articles" 200 "Get articles list"
test_endpoint "POST" "/api/articles" 200 "Create article" \
    '{"article": {"title": "Test Article", "description": "Test desc", "body": "Test body", "tagList": ["test"]}}'
test_endpoint "GET" "/api/articles/test-slug" 200 "Get article by slug"
test_endpoint "PUT" "/api/articles/test-slug" 200 "Update article" \
    '{"article": {"title": "Updated Article"}}'
test_endpoint "DELETE" "/api/articles/test-slug" 200 "Delete article"

# Profile endpoints
echo -e "\n${YELLOW}👥 Profile Endpoints${NC}"
test_endpoint "GET" "/api/profiles/testuser" 200 "Get user profile"
test_endpoint "POST" "/api/profiles/testuser/follow" 200 "Follow user"
test_endpoint "DELETE" "/api/profiles/testuser/follow" 200 "Unfollow user"

# Tags endpoints
echo -e "\n${YELLOW}🏷️  Tags Endpoints${NC}"
test_endpoint "GET" "/api/tags" 200 "Get all tags"

# Comments endpoints (если будут)
echo -e "\n${YELLOW}💬 Comments Endpoints${NC}"
test_endpoint "GET" "/api/articles/test-slug/comments" 200 "Get comments"
test_endpoint "POST" "/api/articles/test-slug/comments" 200 "Add comment" \
    '{"comment": {"body": "Test comment"}}'

# Favorites endpoints
echo -e "\n${YELLOW}❤️  Favorites Endpoints${NC}"
test_endpoint "POST" "/api/articles/test-slug/favorite" 200 "Favorite article"
test_endpoint "DELETE" "/api/articles/test-slug/favorite" 200 "Unfavorite article"

# Проверка несуществующих endpoints
echo -e "\n${YELLOW}🚫 Error Handling${NC}"
test_endpoint "GET" "/api/nonexistent" 404 "Non-existent endpoint"
test_endpoint "GET" "/nonexistent" 404 "Without API prefix"

echo -e "\n${BLUE}📊 Test Summary${NC}"
echo "================================"
echo "✅ Все endpoints должны отвечать статусом 200 (пока пустые handlers)"
echo "❌ 404 ошибки для несуществующих путей - это нормально"
echo ""
echo -e "${YELLOW}💡 Следующие шаги:${NC}"
echo "1. Реализовать логику в handlers (user.go, articles.go, etc.)"
echo "2. Добавить аутентификацию (JWT токены)"
echo "3. Подключить базу данных к handlers"
echo "4. Добавить валидацию входных данных"

echo -e "\n${GREEN}🏁 Testing completed!${NC}"