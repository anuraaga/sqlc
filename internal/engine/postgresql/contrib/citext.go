// Code generated by sqlc-pg-gen. DO NOT EDIT.

package contrib

import (
	"github.com/anuraaga/sqlc/internal/sql/ast"
	"github.com/anuraaga/sqlc/internal/sql/catalog"
)

var funcsCitext = []*catalog.Function{
	{
		Name: "citext",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "boolean"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "citext",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "character"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "citext",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "inet"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "citext_cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "citext_eq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_ge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_gt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_hash",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "citext_hash_extended",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "bigint"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bigint"},
	},
	{
		Name: "citext_larger",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "citext_le",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_lt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_ne",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_pattern_cmp",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "citext_pattern_ge",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_pattern_gt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_pattern_le",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_pattern_lt",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "citext_smaller",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "citextin",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "cstring"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "citextout",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "cstring"},
	},
	{
		Name: "citextsend",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "bytea"},
	},
	{
		Name: "max",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "min",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "citext"},
	},
	{
		Name: "regexp_match",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "regexp_match",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "regexp_matches",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "regexp_matches",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "regexp_replace",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "regexp_replace",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "regexp_split_to_array",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "regexp_split_to_array",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text[]"},
	},
	{
		Name: "regexp_split_to_table",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "regexp_split_to_table",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "replace",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "split_part",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "integer"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
	{
		Name: "strpos",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "integer"},
	},
	{
		Name: "texticlike",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticlike",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticnlike",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticnlike",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticregexeq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticregexeq",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticregexne",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "texticregexne",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "boolean"},
	},
	{
		Name: "translate",
		Args: []*catalog.Argument{
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "citext"},
			},
			{
				Type: &ast.TypeName{Name: "text"},
			},
		},
		ReturnType: &ast.TypeName{Name: "text"},
	},
}

func Citext() *catalog.Schema {
	s := &catalog.Schema{Name: "pg_catalog"}
	s.Funcs = funcsCitext
	return s
}
