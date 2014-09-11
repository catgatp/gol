// (c) 2014 Cergoo
// under terms of ISC license

/*
Package binaryED it's a binary encode/decode implementation
fork github.com/youtube/vitess/go/bson

Attention!

Before you can use this package need to patch standard library reflect,
for unto this add the file the following function:
go/src/pkg/reflect/value.go
    // the Go a user hack
    func (v Value) Ptr() unsafe.Pointer {
       return v.ptr
    }

The package is designed for fast and easy serialization / deserialization:
	uint8, uint16, uint32, uint64, uint
	int8, int16, int32, int64, int
	floate32, floate64
	bool
	string
	slise, array
	map (keys not pointer type)
	struct

Description of the coding format

1. Binary encoding of the data on the basis of binary.LittleEndian.
2. Before the reference types: slices, maps, pointers, added byte: 0-nil, 1-not nil.
3. Before the strings, arrays, slices, maps, added uint32 number of items.
4. Before the data encoded from the source type Interface real name is added to the string type:
in the following format: uint8 length of the string, the string itself is the name of the type.
This is done in order to then be able to decode them by finding the type by name
which must be registered in the decoder prior to decoding.

Important !

- Nonexported field structures are ignored.

- In decoding the variable structure is used in which the decoding occurs,
necessary to match the structure of the receiver structure of the source
up to the order of the fields in the description of the structures.

- The free encoding / decoding only a strictly structured data
custom data types that were present in the sources can type interface is necessary to register the decoder prior to decoding.
That is to map [string] interfase {} in the decoder will need to register custom data types that may be contained in the fields of type interface.

======================================================================
Ru Lang

Пакет binaryED это реализация бинарного кодирования / декодирования на языке Go
fork github.com/youtube/vitess/go/bson

Внимание!

Прежде чем вы сможете использовать этот пакет необходимо пропатчить стандартную библиотеку reflect,
для этого необходимо добавить в файл следующую функцию:
go/src/pkg/reflect/value.go
    // the Go a user hack
    func (v Value) Ptr() unsafe.Pointer {
       return v.ptr
    }

Пакет предназначен для быстрой и лёгкой сериализации / десериализации:
uint8, uint16, uint32, uint64, uint
	int8, int16, int32, int64, int
	floate32, floate64
	bool
	string
	slise, array
	map (keys not pointer type)
	struct

Описание формата кодирования

1. Бинарное кодирование данных на основе binary.LittleEndian.
2. Перед ссылочными типами: срезами, отображениями, указателями, добавляется байт: 0-nil, 1-not nil.
3. Перед строками, массивами, срезами, отображениями, добавляется uint32 количество элементов.
4. Перед данными кодируемыми из источника типа Interface дабавляется строковое реальное наименование типа:
в следующем формате: uint8 длина строки, сама строка наименования типа.
Это делается для того чтобы потом была возможность декодировать их отыскав тип по наименованию,
который должен быть зарегистрирован в декодере перед декодированием.

Важно!

- При кодировании/декодировании неэкспортируемые поля структур игнорируются.

- При декодировании используется структура переменной в которую происходит декодирование, необходимо чтобы структура приёмнника соответсвовала структуре источника вплоть до порядка следования полей в описании структур.

- Возможно свободное кодирование/декодирование только строго структурированных данных,
пользовательские типы данных которые могут присутсвовать в источниках типа interface необходимо регистрировать в декодере перед декодированием.
То есть для map[string]interfase{} в декодере необходимо будет зарегистрировать пользовательские типы данных которые могут содержаться в полях типа interface.
*/
package binaryED
