// Code generated by ragel DO NOT EDIT.
package gir

%%{
  machine graphql_collections;

  include graphql_common "graphql_common.rl";

  # List              = []@{ fcall Value; } ']';
  # ListParser        := List @{ fret; };

  # ObjectField       = Name ':' @{ fcall Value; };
  # Object            = []@{ fcall ObjectField; } '}';
  # ObjectParser      := Object @{ fret; };


  # B3. Query Document

  # Document            =  Definition
  # Definition          = (OperationDefinition | FragmentDefinition)
  # OperationDefinition = (SelectionSet | OperationType Name VariableDefinitions Directives SelectionSet)
  # OperationType         = ('query' | 'mutation');

main := |*
					Variable => { em.Emit(ts, te, Variable, data); };
					IntValue => { em.Emit(ts, te, IntValue, data); };
					FloatValue => { em.Emit(ts, te, FloatValue, data); };
					StringValue => { em.Emit(ts, te, StringValue, data); };
					BooleanValue => { em.Emit(ts, te, BooleanValue, data); };
					NullValue => { em.Emit(ts, te, NullValue, data); };
					EnumValue => { em.Emit(ts, te, EnumValue, data); };
					Ignored => { /* Ignore */ };
				*|;
}%%

%% write data nofinal;

func ParseQuery(data []byte, em Emitter) {
  var token Token
  var stack []int
  cs, p, pe, eof := 0, 0, len(data), len(data)
  ts, te, act, top := 0, 0, 0, 0
  _ = eof
  _ = act
  _ = token
  _ = stack
  _ = top

  %% write init;

  %% write exec;
}
