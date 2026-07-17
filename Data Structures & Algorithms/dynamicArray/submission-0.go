type DynamicArray struct {
	arr      []int
	size     int
	capacity int
}

// NewDynamicArray initializes an empty array with a starting capacity.
func NewDynamicArray(capacity int) *DynamicArray {
	return &DynamicArray{
		arr:      make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

// Get returns the element at index i. Assumes i is valid.
func (this *DynamicArray) Get(i int) int {
	return this.arr[i]
}

// Set sets the element at index i to n. Assumes i is valid.
func (this *DynamicArray) Set(i int, n int) {
	this.arr[i] = n
}

// Pushback pushes the element n to the end of the array.
// Resizes the array first if it is completely full.
func (this *DynamicArray) Pushback(n int) {
	if this.size == this.capacity {
		this.Resize()
	}
	this.arr[this.size] = n
	this.size++
}

// Popback pops and returns the element at the end of the array.
// Assumes the array is non-empty.
func (this *DynamicArray) Popback() int {
	this.size--
	return this.arr[this.size]
}

// Resize doubles the capacity of the array and copies over existing elements.
func (this *DynamicArray) Resize() {
	this.capacity *= 2
	newArr := make([]int, this.capacity)
	
	// Copy elements from the old slice to the new slice
	for i := 0; i < this.size; i++ {
		newArr[i] = this.arr[i]
	}
	this.arr = newArr
}

// GetSize returns the current number of elements in the array.
func (this *DynamicArray) GetSize() int {
	return this.size
}

// GetCapacity returns the current maximum capacity of the array.
func (this *DynamicArray) GetCapacity() int {
	return this.capacity
}