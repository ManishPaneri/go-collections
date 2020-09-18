package binaryheap

import "go-collections/containers"

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Heap)(nil)
	var _ containers.JSONDeserializer = (*Heap)(nil)
}

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
