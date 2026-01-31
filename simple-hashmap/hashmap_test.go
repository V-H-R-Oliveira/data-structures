package main

import (
	"fmt"
	"testing"
)

// Test that empty keys are rejected
func TestPut_EmptyKeyRejected(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Attempt to insert with empty key
	err = hm.Put("", "some value")

	if err == nil {
		t.Error("Expected error when inserting empty key, got nil")
	}

	if hm.occupied != 0 {
		t.Errorf("Expected occupied to be 0 after rejected insertion, got %d", hm.occupied)
	}
}

// Test basic insertion of new entries
func TestPut_NewEntry(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	err = hm.Put("name", "Alice")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if hm.occupied != 1 {
		t.Errorf("Expected occupied to be 1, got %d", hm.occupied)
	}

	value, err := hm.Get("name")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if value != "Alice" {
		t.Errorf("Expected value 'Alice', got '%s'", value)
	}
}

// Test updating existing key (should not increase occupied count)
func TestPut_UpdateExistingKey(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	err = hm.Put("name", "Alice")
	if err != nil {
		t.Fatalf("Unexpected error on first Put: %v", err)
	}

	err = hm.Put("name", "Bob")
	if err != nil {
		t.Fatalf("Unexpected error on update Put: %v", err)
	}

	if hm.occupied != 1 {
		t.Errorf("Expected occupied to be 1 after update, got %d", hm.occupied)
	}

	value, err := hm.Get("name")
	if err != nil {
		t.Fatalf("Unexpected error on Get: %v", err)
	}
	if value != "Bob" {
		t.Errorf("Expected updated value 'Bob', got '%s'", value)
	}
}

// Test multiple different entries
func TestPut_MultipleEntries(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	entries := map[string]string{
		"name":  "Alice",
		"age":   "30",
		"city":  "NYC",
		"email": "alice@example.com",
	}

	for k, v := range entries {
		err := hm.Put(k, v)
		if err != nil {
			t.Fatalf("Unexpected error inserting key '%s': %v", k, err)
		}
	}

	if hm.occupied != 4 {
		t.Errorf("Expected occupied to be 4, got %d", hm.occupied)
	}

	for k, expectedValue := range entries {
		actualValue, err := hm.Get(k)
		if err != nil {
			t.Errorf("Unexpected error getting key '%s': %v", k, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", k, expectedValue, actualValue)
		}
	}
}

// Test collision handling - insert keys that hash to same bucket
func TestPut_CollisionHandling(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Generate keys that will collide
	// We'll insert multiple keys and verify they're all retrievable
	keys := []string{"key1", "key2", "key3", "key4", "key5"}

	for i, key := range keys {
		err := hm.Put(key, fmt.Sprintf("value%d", i))
		if err != nil {
			t.Fatalf("Unexpected error inserting key '%s': %v", key, err)
		}
	}

	// Verify all keys are retrievable
	for i, key := range keys {
		expectedValue := fmt.Sprintf("value%d", i)
		actualValue, err := hm.Get(key)
		if err != nil {
			t.Errorf("Unexpected error getting key '%s': %v", key, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}
}

// Test rehashing when load factor exceeds 0.75
func TestPut_RehashingOnLoadFactor(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	initialSize := hm.size
	// Insert enough entries to trigger rehash
	// 64 * 0.75 = 48, so inserting 49 should trigger rehash
	numEntries := uint64(49)

	for i := range numEntries {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := hm.Put(key, value)
		if err != nil {
			t.Fatalf("Unexpected error inserting key '%s': %v", key, err)
		}
	}

	if hm.occupied != numEntries {
		t.Errorf("Expected occupied to be %d, got %d", numEntries, hm.occupied)
	}

	// Check if size increased (rehashed)
	if hm.size <= initialSize {
		t.Errorf("Expected size to increase after rehashing. Initial: %d, Current: %d", initialSize, hm.size)
	}

	// Verify all entries are still retrievable after rehash
	for i := range numEntries {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		actualValue, err := hm.Get(key)
		if err != nil {
			t.Errorf("Unexpected error getting key '%s' after rehash: %v", key, err)
		}
		if actualValue != expectedValue {
			t.Errorf("After rehash - Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}

	// Expected new size should be double the original
	expectedSize := initialSize * 2
	if hm.size != expectedSize {
		t.Errorf("Expected size to be %d after rehashing, got %d", expectedSize, hm.size)
	}
}

// Test multiple rehashing cycles
func TestPut_MultipleRehashes(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert enough to trigger multiple rehashes
	// Start: 64, after 48: 128, after 96: 256
	numEntries := uint64(150)

	for i := range numEntries {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := hm.Put(key, value)
		if err != nil {
			t.Fatalf("Unexpected error inserting key '%s': %v", key, err)
		}
	}

	if hm.occupied != numEntries {
		t.Errorf("Expected occupied to be %d, got %d", numEntries, hm.occupied)
	}

	// Should have rehashed at least twice
	if hm.size < 256 {
		t.Errorf("Expected size to be at least 256 after multiple rehashes, got %d", hm.size)
	}

	// Verify all entries are still correct
	for i := range numEntries {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		actualValue, err := hm.Get(key)
		if err != nil {
			t.Errorf("Unexpected error getting key '%s' after multiple rehashes: %v", key, err)
		}
		if actualValue != expectedValue {
			t.Errorf("After multiple rehashes - Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}
}

// Test updating keys after rehashing
func TestPut_UpdateAfterRehash(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert enough to trigger rehash
	numEntries := uint64(50)

	for i := range numEntries {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := hm.Put(key, value)
		if err != nil {
			t.Fatalf("Unexpected error inserting key '%s': %v", key, err)
		}
	}

	// Now update some existing keys
	err = hm.Put("key0", "updated0")
	if err != nil {
		t.Fatalf("Unexpected error updating key0: %v", err)
	}

	err = hm.Put("key25", "updated25")
	if err != nil {
		t.Fatalf("Unexpected error updating key25: %v", err)
	}

	if hm.occupied != numEntries {
		t.Errorf("Expected occupied to remain %d after updates, got %d", numEntries, hm.occupied)
	}

	val0, err := hm.Get("key0")
	if err != nil {
		t.Fatalf("Unexpected error getting key0: %v", err)
	}
	if val0 != "updated0" {
		t.Errorf("Expected updated value 'updated0', got '%s'", val0)
	}

	val25, err := hm.Get("key25")
	if err != nil {
		t.Fatalf("Unexpected error getting key25: %v", err)
	}
	if val25 != "updated25" {
		t.Errorf("Expected updated value 'updated25', got '%s'", val25)
	}
}

// Test exact load factor threshold (edge case)
func TestPut_ExactLoadFactorThreshold(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert exactly to the threshold: 64 * 0.75 = 48
	for i := 0; i < 48; i++ {
		key := fmt.Sprintf("key%d", i)
		err := hm.Put(key, "value")
		if err != nil {
			t.Fatalf("Unexpected error inserting key '%s': %v", key, err)
		}
	}

	// At this point, should be at exactly 0.75 load factor
	currentLoadFactor := float64(hm.occupied) / float64(hm.size)
	if currentLoadFactor != 0.75 {
		t.Errorf("Expected load factor to be 0.75, got %f", currentLoadFactor)
	}

	// One more insert should trigger rehash
	initialSize := hm.size
	err = hm.Put("trigger", "rehash")
	if err != nil {
		t.Fatalf("Unexpected error on trigger insert: %v", err)
	}

	if hm.size <= initialSize {
		t.Errorf("Expected rehash after exceeding threshold")
	}
}

// Test that empty key doesn't affect occupied count
func TestPut_EmptyKeyDoesNotAffectOccupied(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert valid entries
	err = hm.Put("key1", "value1")
	if err != nil {
		t.Fatalf("Unexpected error inserting key1: %v", err)
	}

	err = hm.Put("key2", "value2")
	if err != nil {
		t.Fatalf("Unexpected error inserting key2: %v", err)
	}

	occupiedBefore := hm.occupied

	// Try to insert empty key
	err = hm.Put("", "should fail")
	if err == nil {
		t.Error("Expected error for empty key")
	}

	if hm.occupied != occupiedBefore {
		t.Errorf("Expected occupied to remain %d, got %d", occupiedBefore, hm.occupied)
	}
}

// Test updating with empty key is also rejected
func TestPut_UpdateWithEmptyKeyRejected(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert a valid entry first
	err = hm.Put("validKey", "originalValue")
	if err != nil {
		t.Fatalf("Unexpected error inserting validKey: %v", err)
	}

	// Try to update with empty key
	err = hm.Put("", "newValue")

	if err == nil {
		t.Error("Expected error when updating with empty key")
	}

	// Original entry should be unaffected
	val, err := hm.Get("validKey")
	if err != nil {
		t.Fatalf("Unexpected error getting validKey: %v", err)
	}
	if val != "originalValue" {
		t.Errorf("Expected original value to be unchanged, got '%s'", val)
	}
}

// Test creating HashMap with invalid initial size
func TestNewHashMap_InvalidSize(t *testing.T) {
	testCases := []struct {
		name        string
		initialSize uint64
	}{
		{"zero size", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			hm, err := NewHashMap(tc.initialSize)
			if err == nil {
				t.Errorf("Expected error for initial size %d, got nil", tc.initialSize)
			}
			if hm != nil {
				t.Errorf("Expected nil HashMap for invalid size, got non-nil")
			}
		})
	}
}

// Test creating HashMap with valid initial size
func TestNewHashMap_ValidSize(t *testing.T) {
	testCases := []uint64{1, 16, 64, 128, 1024}

	for _, size := range testCases {
		t.Run(fmt.Sprintf("size_%d", size), func(t *testing.T) {
			hm, err := NewHashMap(size)
			if err != nil {
				t.Errorf("Unexpected error for size %d: %v", size, err)
			}
			if hm == nil {
				t.Errorf("Expected non-nil HashMap for size %d", size)
			}
			if hm != nil && hm.size != size {
				t.Errorf("Expected size %d, got %d", size, hm.size)
			}
		})
	}
}

// ============================================
// GET FUNCTION TESTS
// ============================================

// Test getting a key that doesn't exist - should return empty string and nil error
func TestGet_NonExistentKey(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	value, err := hm.Get("nonexistent")
	if err != nil {
		t.Errorf("Expected nil error when getting non-existent key, got: %v", err)
	}
	if value != "" {
		t.Errorf("Expected empty value for non-existent key, got '%s'", value)
	}
}

// Test getting with empty key
func TestGet_EmptyKey(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	value, err := hm.Get("")
	if err == nil {
		t.Error("Expected error when getting with empty key, got nil")
	}
	if value != "" {
		t.Errorf("Expected empty value for empty key, got '%s'", value)
	}
}

// Test getting a single existing key
func TestGet_ExistingKey(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	err = hm.Put("name", "Alice")
	if err != nil {
		t.Fatalf("Failed to put key: %v", err)
	}

	value, err := hm.Get("name")
	if err != nil {
		t.Errorf("Unexpected error getting existing key: %v", err)
	}
	if value != "Alice" {
		t.Errorf("Expected 'Alice', got '%s'", value)
	}
}

// Test getting multiple keys
func TestGet_MultipleKeys(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	testData := map[string]string{
		"name":    "Alice",
		"age":     "30",
		"city":    "NYC",
		"country": "USA",
		"job":     "Engineer",
	}

	// Insert all keys
	for k, v := range testData {
		err := hm.Put(k, v)
		if err != nil {
			t.Fatalf("Failed to put key '%s': %v", k, err)
		}
	}

	// Get and verify all keys
	for k, expectedValue := range testData {
		actualValue, err := hm.Get(k)
		if err != nil {
			t.Errorf("Unexpected error getting key '%s': %v", k, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", k, expectedValue, actualValue)
		}
	}
}

// Test getting a key after it has been updated
func TestGet_AfterUpdate(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert initial value
	err = hm.Put("status", "pending")
	if err != nil {
		t.Fatalf("Failed to put initial value: %v", err)
	}

	// Verify initial value
	value, err := hm.Get("status")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if value != "pending" {
		t.Errorf("Expected 'pending', got '%s'", value)
	}

	// Update value
	err = hm.Put("status", "completed")
	if err != nil {
		t.Fatalf("Failed to update value: %v", err)
	}

	// Verify updated value
	value, err = hm.Get("status")
	if err != nil {
		t.Errorf("Unexpected error after update: %v", err)
	}
	if value != "completed" {
		t.Errorf("Expected 'completed', got '%s'", value)
	}
}

// Test getting keys that caused collisions
func TestGet_CollidedKeys(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert multiple keys that will likely cause collisions
	numKeys := 100
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := hm.Put(key, value)
		if err != nil {
			t.Fatalf("Failed to put key '%s': %v", key, err)
		}
	}

	// Verify all keys are retrievable
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		actualValue, err := hm.Get(key)
		if err != nil {
			t.Errorf("Failed to get key '%s': %v", key, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}
}

// Test getting keys after rehashing
func TestGet_AfterRehash(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert enough keys to trigger rehash
	numKeys := 50
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := hm.Put(key, value)
		if err != nil {
			t.Fatalf("Failed to put key '%s': %v", key, err)
		}
	}

	// Verify table was rehashed
	if hm.size <= 64 {
		t.Errorf("Expected rehash to occur, size is still %d", hm.size)
	}

	// Verify all keys are still retrievable after rehash
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		actualValue, err := hm.Get(key)
		if err != nil {
			t.Errorf("Failed to get key '%s' after rehash: %v", key, err)
		}
		if actualValue != expectedValue {
			t.Errorf("After rehash - Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}
}

// Test getting keys after multiple rehashes
func TestGet_AfterMultipleRehashes(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert enough to trigger multiple rehashes
	numKeys := 200
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		value := fmt.Sprintf("value%d", i)
		err := hm.Put(key, value)
		if err != nil {
			t.Fatalf("Failed to put key '%s': %v", key, err)
		}
	}

	// Verify multiple rehashes occurred
	if hm.size < 256 {
		t.Errorf("Expected multiple rehashes, size is %d", hm.size)
	}

	// Verify all keys are still retrievable
	for i := 0; i < numKeys; i++ {
		key := fmt.Sprintf("key%d", i)
		expectedValue := fmt.Sprintf("value%d", i)
		actualValue, err := hm.Get(key)
		if err != nil {
			t.Errorf("Failed to get key '%s' after multiple rehashes: %v", key, err)
		}
		if actualValue != expectedValue {
			t.Errorf("After multiple rehashes - Key '%s': expected '%s', got '%s'", key, expectedValue, actualValue)
		}
	}
}

// Test getting a mix of existing and non-existing keys
func TestGet_MixedKeys(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert some keys
	existingKeys := []string{"key1", "key3", "key5", "key7"}
	for _, key := range existingKeys {
		err := hm.Put(key, "value_"+key)
		if err != nil {
			t.Fatalf("Failed to put key '%s': %v", key, err)
		}
	}

	// Test existing keys
	for _, key := range existingKeys {
		value, err := hm.Get(key)
		if err != nil {
			t.Errorf("Unexpected error getting existing key '%s': %v", key, err)
		}
		expectedValue := "value_" + key
		if value != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", key, expectedValue, value)
		}
	}

	// Test non-existing keys - should return empty string with nil error
	nonExistingKeys := []string{"key2", "key4", "key6", "key8"}
	for _, key := range nonExistingKeys {
		value, err := hm.Get(key)
		if err != nil {
			t.Errorf("Expected nil error for non-existing key '%s', got: %v", key, err)
		}
		if value != "" {
			t.Errorf("Expected empty value for non-existing key '%s', got '%s'", key, value)
		}
	}
}

// Test getting with special characters in keys
func TestGet_SpecialCharacterKeys(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	specialKeys := map[string]string{
		"key-with-dash":       "value1",
		"key_with_underscore": "value2",
		"key.with.dot":        "value3",
		"key@with@at":         "value4",
		"key#with#hash":       "value5",
		"key with space":      "value6",
	}

	// Insert all special keys
	for k, v := range specialKeys {
		err := hm.Put(k, v)
		if err != nil {
			t.Fatalf("Failed to put key '%s': %v", k, err)
		}
	}

	// Get and verify all special keys
	for k, expectedValue := range specialKeys {
		actualValue, err := hm.Get(k)
		if err != nil {
			t.Errorf("Failed to get key '%s': %v", k, err)
		}
		if actualValue != expectedValue {
			t.Errorf("Key '%s': expected '%s', got '%s'", k, expectedValue, actualValue)
		}
	}
}

// Test getting empty values
func TestGet_EmptyValues(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Put keys with empty values
	err = hm.Put("emptyValue", "")
	if err != nil {
		t.Fatalf("Failed to put key with empty value: %v", err)
	}

	// Get should return empty string without error
	value, err := hm.Get("emptyValue")
	if err != nil {
		t.Errorf("Unexpected error getting key with empty value: %v", err)
	}
	if value != "" {
		t.Errorf("Expected empty value, got '%s'", value)
	}
}

// Test case sensitivity of keys
func TestGet_CaseSensitiveKeys(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert keys with different cases
	err = hm.Put("Name", "Alice")
	if err != nil {
		t.Fatalf("Failed to put 'Name': %v", err)
	}

	err = hm.Put("name", "Bob")
	if err != nil {
		t.Fatalf("Failed to put 'name': %v", err)
	}

	err = hm.Put("NAME", "Charlie")
	if err != nil {
		t.Fatalf("Failed to put 'NAME': %v", err)
	}

	// Verify each key returns its own value
	value1, err := hm.Get("Name")
	if err != nil {
		t.Errorf("Failed to get 'Name': %v", err)
	}
	if value1 != "Alice" {
		t.Errorf("Expected 'Alice' for 'Name', got '%s'", value1)
	}

	value2, err := hm.Get("name")
	if err != nil {
		t.Errorf("Failed to get 'name': %v", err)
	}
	if value2 != "Bob" {
		t.Errorf("Expected 'Bob' for 'name', got '%s'", value2)
	}

	value3, err := hm.Get("NAME")
	if err != nil {
		t.Errorf("Failed to get 'NAME': %v", err)
	}
	if value3 != "Charlie" {
		t.Errorf("Expected 'Charlie' for 'NAME', got '%s'", value3)
	}
}

// Test getting non-existent key from empty hashmap
func TestGet_EmptyHashMap(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	value, err := hm.Get("anykey")
	if err != nil {
		t.Errorf("Expected nil error for non-existent key in empty map, got: %v", err)
	}
	if value != "" {
		t.Errorf("Expected empty value for non-existent key in empty map, got '%s'", value)
	}
}

// Test getting non-existent keys after some insertions and deletions (if delete is implemented)
func TestGet_NonExistentAfterOperations(t *testing.T) {
	hm, err := NewHashMap(64)
	if err != nil {
		t.Fatalf("Failed to create HashMap: %v", err)
	}

	// Insert some keys
	hm.Put("key1", "value1")
	hm.Put("key2", "value2")
	hm.Put("key3", "value3")

	// Try to get keys that were never inserted
	value, err := hm.Get("key4")
	if err != nil {
		t.Errorf("Expected nil error for non-existent key 'key4', got: %v", err)
	}
	if value != "" {
		t.Errorf("Expected empty value for non-existent key 'key4', got '%s'", value)
	}

	value, err = hm.Get("nonexistent")
	if err != nil {
		t.Errorf("Expected nil error for non-existent key 'nonexistent', got: %v", err)
	}
	if value != "" {
		t.Errorf("Expected empty value for non-existent key 'nonexistent', got '%s'", value)
	}
}
