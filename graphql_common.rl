%%{
  machine graphql_common;
Punctuator        = ('!' | '$' | '(' | ')' | '...' | ':' | '=' |
                        '@' | '[' | ']' | '{' | '|' | '}');

  SourceCharacter   = 0x09 | 0x0a | 0x0d | (0x20..0x7e) |
                      (0xc2..0xdf) (0x80..0xbf) |
                      (0xe0..0xef 0x80..0xbf 0x80..0xbf) |
                      (0xf0..0xf4 0x80..0xbf 0x80..0xbf 0x80..0xbf);
  
  # B1. Ignored Tokens

  UnicodeBOM        = (0xfe 0xff);
  WhiteSpace        = 0x09 | 0x20;
  LineTerminator    = 0x0a | (0x0d 0x0a) | 0x0d;
  Comment           = '#'(SourceCharacter -- LineTerminator)*;
  Comma             = ',';
  Ignored           = (UnicodeBOM | WhiteSpace | LineTerminator | Comment | Comma);

  # B2. Tokens

  EscapedUnicode    = [0-9A-Fa-f]{4};

  EscapedCharacter  = ["\\bfnrt];

  Name              = [_A-Za-z][_0-9A-Za-z]*;

  FloatValue        = '-'?[0-9]*"."[0-9]*([eE][\-+]?[0-9]+)?;

  IntValue          = '-'?[1-9][0-9]*;

  StringValue       = '"'(SourceCharacter -- ('\\' | '\"') |
                      '\\u'EscapedUnicode |
                      '\\'EscapedCharacter)*'"';

  BooleanValue      = ('true' | 'false');

  NullValue         = 'null';

  Variable          = '$'Name;

  EnumValue         = (Name -- (BooleanValue | NullValue));

  PrimitiveValue    = (Variable | IntValue | FloatValue | StringValue |
                       BooleanValue | NullValue | EnumValue);


  #ListStart         = '[';
  #List              = (PrimitiveValue)* ']';
  #ListParser       := List @{ fret; };
}%%
