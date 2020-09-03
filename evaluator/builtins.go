package evaluator

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"github.com/emilkloeden/monkey/object"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return NULL
		},
	},
	"exit": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			os.Exit(0)
			return NULL
		},
	},
	"join": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `join` must be ARRAY, got %s",
					args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `join` must be STRING, got %s",
					args[1].Type())
			}

			array, ok := args[0].(*object.Array)
			if !ok {
				return newError("first argument to `join` must be ARRAY, got=%T(%+v)", args[0], args[0])
			}
			arrayElements := array.Elements
			var out bytes.Buffer
			length := len(arrayElements)

			if length == 0 {
				return &object.String{Value: ""}
			}

			if length == 1 {
				return &object.String{Value: arrayElements[0].Inspect()}
			}

			max := length - 1

			for idx, el := range arrayElements {
				out.WriteString(el.Inspect())
				if idx < max {
					out.WriteString(args[1].Inspect())
				}
			}
			return &object.String{Value: out.String()}
		},
	},
	"split": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}

			str, ok := args[0].(*object.String)
			if !ok {
				return newError("first argument to `split` must be STRING, got=%T(%+v)", args[0], args[0])
			}

			delim, ok := args[1].(*object.String)
			if !ok {
				return newError("second argument to `split` must be STRING, got=%T(%+v)", args[0], args[0])
			}

			elements := strings.Split(str.Value, delim.Value)
			strElements := make([]object.Object, 0)

			for _, el := range elements {
				strElements = append(strElements, &object.String{Value: el})
			}

			return &object.Array{Elements: strElements}
		},
	},
	"keys": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var hash *object.Hash
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			switch args[0].Type() {
			case object.HASH_OBJ:
				hash = args[0].(*object.Hash)
			case object.MODULE:
				module, ok := args[0].(*object.Module)
				if !ok {
					return newError("argument to `keys` must be HASH or MODULE, got=%T(%+v)", args[0], args[0])
				}
				hash = module.Attrs.(*object.Hash)
			default:
				return newError("argument to `keys` must be HASH or MODULE, got=%T(%+v)", args[0], args[0])
			}

			keys := make([]object.Object, 0, len(hash.Pairs))
			for _, v := range hash.Pairs {
				keys = append(keys, v.Key)
			}

			return &object.Array{Elements: keys}
		},
	},
	"values": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var hash *object.Hash
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch args[0].Type() {
			case object.HASH_OBJ:
				hash = args[0].(*object.Hash)
			case object.MODULE:
				module, ok := args[0].(*object.Module)
				if !ok {
					return newError("argument to `keys` must be HASH or MODULE, got=%T(%+v)", args[0], args[0])
				}
				hash = module.Attrs.(*object.Hash)
			default:
				return newError("argument to `keys` must be HASH or MODULE, got=%T(%+v)", args[0], args[0])
			}

			values := make([]object.Object, 0, len(hash.Pairs))
			for _, v := range hash.Pairs {
				values = append(values, v.Value)
			}

			return &object.Array{Elements: values}
		},
	},
}
