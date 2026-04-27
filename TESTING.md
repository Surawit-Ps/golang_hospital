# Unit Testing Guide

This document explains how to run and manage unit tests for the Hospital Management System API.

## Setup

### Install Test Dependencies

The project uses `testify` for mocking and assertions. Install it:

```bash
cd backend
go get github.com/stretchr/testify
```

## Running Tests

### Run All Tests

```bash
# From backend directory
cd backend

# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with coverage
go test -v -cover ./...
```

### Run Specific Test File

```bash
# Test hospital handlers
go test -v ./adapter/handler -run Hospital

# Test staff handlers
go test -v ./adapter/handler -run Staff

# Test patient handlers
go test -v ./adapter/handler -run Patient
```

### Run Specific Test

```bash
# Run a single test
go test -v ./adapter/handler -run TestRegisterHospital_Success
```

## Test Coverage

### Generate Coverage Report

```bash
# Generate coverage report
go test -cover ./...

# Generate detailed coverage report as HTML
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### View Coverage in Browser

After generating the HTML report:
- Open `coverage.html` in your browser
- Green areas: covered code
- Red areas: uncovered code

## Test Structure

### Handler Tests (`adapter/handler/*_test.go`)

Tests for HTTP endpoints:
- Request/response handling
- Authorization checks
- Invalid input validation
- Error handling

**Mock Used**: `MockHospitalService`, `MockStaffService`, `MockPatientService`

**Files**:
- `hospital_handler_test.go` - Tests for `/hospitals` endpoints
- `staff_handler_test.go` - Tests for `/staff` endpoints
- `patients_handler_test.go` - Tests for `/patients` endpoints

### Service Tests (`core/service/*_test.go`)

Tests for business logic:
- Data validation
- Service operations
- Error scenarios

**Mock Used**: `MockHospitalRepository`, `MockStaffRepository`, `MockPatientRepository`

**Files**:
- `hospital_service_test.go` - Hospital service logic tests
- `staff_service_test.go` - Staff service logic tests
- `patiients_service_test.go` - Patient service logic tests

## Test Examples

### Example 1: Handler Test

```go
func TestRegisterHospital_Success(t *testing.T) {
    mockService := new(MockHospitalService)
    
    newHospital := &entity.Hospital{
        Id:   "HOSP0001",
        Name: "Hospital Test",
    }
    
    mockService.On("RegisterHospital", newHospital).Return(newHospital, nil)
    
    handler := NewHospitalHandler(mockService)
    
    body, _ := json.Marshal(newHospital)
    req := httptest.NewRequest(http.MethodPost, "/hospitals", bytes.NewBuffer(body))
    req.Header.Set("Content-Type", "application/json")
    
    w := httptest.NewRecorder()
    c := createTestContext(w, req)
    handler.RegisterHospital(c)
    
    assert.Equal(t, http.StatusCreated, w.Code)
    mockService.AssertExpectations(t)
}
```

### Example 2: Service Test

```go
func TestHospitalService_RegisterHospital_Success(t *testing.T) {
    mockRepo := new(MockHospitalRepository)
    
    hospital := &entity.Hospital{
        Name: "Hospital A",
    }
    
    resultHospital := &entity.Hospital{
        Id:   "HOSP0001",
        Name: "Hospital A",
    }
    
    mockRepo.On("CreateHospital", hospital).Return(resultHospital, nil)
    
    service := NewHospitalService(mockRepo)
    result, err := service.RegisterHospital(hospital)
    
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "HOSP0001", result.Id)
    mockRepo.AssertExpectations(t)
}
```

## Test Cases Summary

### Hospital Handler Tests (8 tests)
- ✅ Register Hospital - Success
- ✅ Register Hospital - Invalid JSON
- ✅ Get Hospital Info - Success
- ✅ Get Hospital Info - Not Found
- ✅ Get All Hospitals - Success
- ✅ Get All Hospitals - Empty List
- ✅ Get All Hospitals - Database Error

### Staff Handler Tests (7 tests)
- ✅ Register Staff - Success
- ✅ Register Staff - Unauthorized
- ✅ Register Staff - Invalid JSON
- ✅ Login Staff - Success
- ✅ Login Staff - Invalid Credentials
- ✅ Get User Login - Success
- ✅ Get User Login - Unauthorized
- ✅ Get User Login - Not Found

### Patient Handler Tests (8 tests)
- ✅ Register Patient - Success
- ✅ Register Patient - Unauthorized
- ✅ Register Patient - Invalid JSON
- ✅ Get Patient Info - Success
- ✅ Get Patient Info - Not Found
- ✅ Get Patients - Success (with default pagination)
- ✅ Get Patients - Unauthorized
- ✅ Get Patients - Invalid Page Number
- ✅ Get Patients - Invalid Limit Number
- ✅ Get Patients - Empty List
- ✅ Get Patients - With Filter

### Hospital Service Tests (7 tests)
- ✅ Register Hospital - Success
- ✅ Register Hospital - Database Error
- ✅ Get Hospital Info - Success
- ✅ Get Hospital Info - Not Found
- ✅ Get All Hospitals - Success
- ✅ Get All Hospitals - Empty List
- ✅ Get All Hospitals - Database Error

### Staff Service Tests (7 tests)
- ✅ Register Staff - Success
- ✅ Register Staff - Database Error
- ✅ Login Staff - Success
- ✅ Login Staff - Invalid Credentials
- ✅ Login Staff - User Not Found
- ✅ Get User - Success
- ✅ Get User - Not Found

### Patient Service Tests (8 tests)
- ✅ Register Patient - Success
- ✅ Register Patient - Database Error
- ✅ Get Patient Info - Success
- ✅ Get Patient Info - Not Found
- ✅ Get Patients - Success
- ✅ Get Patients - With Firstname Filter
- ✅ Get Patients - Empty Result
- ✅ Get Patients - Database Error
- ✅ Get Patients - With Pagination

## Continuous Integration

### Run Tests with Make

```bash
# From project root
make -f Makefile test
```

Add to `Makefile`:

```makefile
test:
	cd backend && go test -v -cover ./...

test-coverage:
	cd backend && go test -coverprofile=coverage.out ./...
	cd backend && go tool cover -html=coverage.out -o coverage.html

test-verbose:
	cd backend && go test -v ./...
```

## Best Practices

1. **Test Naming**: Use `Test[Function]_[Scenario]` format
2. **Mock Dependencies**: Always mock external dependencies
3. **Arrange-Act-Assert**: Organize tests in this pattern
4. **Test Isolation**: Each test should be independent
5. **Edge Cases**: Test error conditions and edge cases
6. **Coverage**: Aim for >80% code coverage

## Troubleshooting

### Test Fails with "undefined" error

Make sure all imports are included:

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)
```

### Mock Not Working

Ensure mock is created before calling functions:

```go
mockService := new(MockHospitalService)
mockService.On("RegisterHospital", hospital).Return(hospital, nil)
// Then use the mock
```

### Tests Pass Locally but Fail in CI

- Check environment variables
- Verify database connectivity in CI
- Ensure Go version compatibility

## Integration with Docker

Tests can run in Docker:

```bash
# Build test image
docker build -t hospital-tests -f Dockerfile.test .

# Run tests in container
docker run hospital-tests
```

## Resources

- [Testify Documentation](https://github.com/stretchr/testify)
- [Go Testing Package](https://golang.org/pkg/testing/)
- [HTTP Test Utilities](https://golang.org/pkg/net/http/httptest/)
