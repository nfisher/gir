
//line graphql_collections.rl:1
package gir


//line graphql_collections.rl:33



//line graphql_collections.go:11
const graphql_collections_start int = 19
const graphql_collections_first_final int = 19
const graphql_collections_error int = 0

const graphql_collections_en_main int = 19


//line graphql_collections.rl:36

func ParseQuery(data []byte) {
}

func ParseValues(data []byte, em Emitter) {
  var token Token
  cs, p, pe, eof := 0, 0, len(data), len(data)
  ts, te, act := 0, 0, 0
  _ = eof
  _ = act
  _ = token

  
//line graphql_collections.go:33
	{
	cs = graphql_collections_start
	ts = 0
	te = 0
	act = 0
	}

//line graphql_collections.rl:49

  
//line graphql_collections.go:44
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 19:
		goto st_case_19
	case 0:
		goto st_case_0
	case 20:
		goto st_case_20
	case 1:
		goto st_case_1
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 21:
		goto st_case_21
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 22:
		goto st_case_22
	case 14:
		goto st_case_14
	case 23:
		goto st_case_23
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 24:
		goto st_case_24
	case 17:
		goto st_case_17
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 18:
		goto st_case_18
	}
	goto st_out
tr2:
//line graphql_collections.rl:27
te = p+1
{ em.Emit(ts, te, StringValue, data); }
	goto st19
tr11:
//line graphql_collections.rl:31
p = (te) - 1
{ /* Ignore */ }
	goto st19
tr19:
//line graphql_collections.rl:26
p = (te) - 1
{ em.Emit(ts, te, FloatValue, data); }
	goto st19
tr22:
//line graphql_collections.rl:31
te = p+1
{ /* Ignore */ }
	goto st19
tr31:
//line graphql_collections.rl:31
te = p
p--
{ /* Ignore */ }
	goto st19
tr33:
//line graphql_collections.rl:24
te = p
p--
{ em.Emit(ts, te, Variable, data); }
	goto st19
tr34:
//line graphql_collections.rl:26
te = p
p--
{ em.Emit(ts, te, FloatValue, data); }
	goto st19
tr36:
//line graphql_collections.rl:25
te = p
p--
{ em.Emit(ts, te, IntValue, data); }
	goto st19
tr37:
//line graphql_collections.rl:30
te = p
p--
{ em.Emit(ts, te, EnumValue, data); }
	goto st19
tr50:
//line graphql_collections.rl:28
te = p+1
{ em.Emit(ts, te, BooleanValue, data); }
	goto st19
tr53:
//line graphql_collections.rl:29
te = p+1
{ em.Emit(ts, te, NullValue, data); }
	goto st19
	st19:
//line NONE:1
ts = 0

		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line NONE:1
ts = p

//line graphql_collections.go:213
		switch data[p] {
		case 13:
			goto st20
		case 32:
			goto tr22
		case 34:
			goto st1
		case 35:
			goto tr12
		case 36:
			goto st13
		case 44:
			goto tr22
		case 45:
			goto st14
		case 46:
			goto tr16
		case 48:
			goto st17
		case 95:
			goto st26
		case 102:
			goto st36
		case 110:
			goto st40
		case 116:
			goto st43
		case 254:
			goto st18
		}
		switch {
		case data[p] < 49:
			if 9 <= data[p] && data[p] <= 10 {
				goto tr22
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st26
				}
			case data[p] >= 65:
				goto st26
			}
		default:
			goto st25
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
		if data[p] == 10 {
			goto tr22
		}
		goto tr31
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 13:
			goto st1
		case 34:
			goto tr2
		case 92:
			goto st2
		}
		switch {
		case data[p] < 194:
			switch {
			case data[p] > 10:
				if 32 <= data[p] && data[p] <= 126 {
					goto st1
				}
			case data[p] >= 9:
				goto st1
			}
		case data[p] > 223:
			switch {
			case data[p] > 239:
				if 240 <= data[p] && data[p] <= 244 {
					goto st9
				}
			case data[p] >= 224:
				goto st8
			}
		default:
			goto st7
		}
		goto st0
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
		switch data[p] {
		case 34:
			goto st1
		case 92:
			goto st1
		case 98:
			goto st1
		case 102:
			goto st1
		case 110:
			goto st1
		case 114:
			goto st1
		case 116:
			goto st1
		case 117:
			goto st3
		}
		goto st0
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st4
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st4
			}
		default:
			goto st4
		}
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st5
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st5
			}
		default:
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st6
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st6
			}
		default:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st1
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto st1
			}
		default:
			goto st1
		}
		goto st0
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
		if 128 <= data[p] && data[p] <= 191 {
			goto st1
		}
		goto st0
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
		if 128 <= data[p] && data[p] <= 191 {
			goto st7
		}
		goto st0
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
		if 128 <= data[p] && data[p] <= 191 {
			goto st8
		}
		goto st0
tr12:
//line NONE:1
te = p+1

	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line graphql_collections.go:444
		if data[p] == 9 {
			goto tr12
		}
		switch {
		case data[p] < 194:
			if 32 <= data[p] && data[p] <= 126 {
				goto tr12
			}
		case data[p] > 223:
			switch {
			case data[p] > 239:
				if 240 <= data[p] && data[p] <= 244 {
					goto st12
				}
			case data[p] >= 224:
				goto st11
			}
		default:
			goto st10
		}
		goto tr31
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
		if 128 <= data[p] && data[p] <= 191 {
			goto tr12
		}
		goto tr11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		if 128 <= data[p] && data[p] <= 191 {
			goto st10
		}
		goto tr11
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		if 128 <= data[p] && data[p] <= 191 {
			goto st11
		}
		goto tr11
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		if data[p] == 95 {
			goto st22
		}
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		case data[p] >= 65:
			goto st22
		}
		goto st0
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
		if data[p] == 95 {
			goto st22
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st22
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st22
			}
		default:
			goto st22
		}
		goto tr33
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
		switch data[p] {
		case 46:
			goto tr16
		case 48:
			goto st17
		}
		if 49 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto st0
tr16:
//line NONE:1
te = p+1

	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line graphql_collections.go:556
		switch data[p] {
		case 69:
			goto st15
		case 101:
			goto st15
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto tr16
		}
		goto tr34
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
		switch data[p] {
		case 43:
			goto st16
		case 45:
			goto st16
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st24
		}
		goto tr19
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
		if 48 <= data[p] && data[p] <= 57 {
			goto st24
		}
		goto tr19
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		if 48 <= data[p] && data[p] <= 57 {
			goto st24
		}
		goto tr34
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
		if data[p] == 46 {
			goto tr16
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st17
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 46 {
			goto tr16
		}
		if 48 <= data[p] && data[p] <= 57 {
			goto st25
		}
		goto tr36
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 95:
			goto st26
		case 97:
			goto st28
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 108:
			goto st29
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 115:
			goto st33
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		case 117:
			goto st34
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 114:
			goto st32
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		case 117:
			goto st33
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			switch {
			case data[p] > 100:
				if 103 <= data[p] && data[p] <= 122 {
					goto st26
				}
			case data[p] >= 97:
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 108:
			goto st35
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			switch {
			case data[p] > 107:
				if 109 <= data[p] && data[p] <= 122 {
					goto st26
				}
			case data[p] >= 97:
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
		switch data[p] {
		case 95:
			goto st26
		case 97:
			goto st37
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 98 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 108:
			goto st38
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 115:
			goto st39
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
		switch data[p] {
		case 95:
			goto st26
		case 101:
			goto tr50
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		case 117:
			goto st41
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 108:
			goto st42
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 108:
			goto tr53
		case 110:
			goto st30
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 114:
			goto st44
		case 116:
			goto st31
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
		switch data[p] {
		case 95:
			goto st26
		case 102:
			goto st27
		case 110:
			goto st30
		case 116:
			goto st31
		case 117:
			goto st39
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto st26
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto st26
			}
		default:
			goto st26
		}
		goto tr37
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
		if data[p] == 255 {
			goto tr22
		}
		goto st0
	st_out:
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 20:
			goto tr31
		case 21:
			goto tr31
		case 10:
			goto tr11
		case 11:
			goto tr11
		case 12:
			goto tr11
		case 22:
			goto tr33
		case 23:
			goto tr34
		case 15:
			goto tr19
		case 16:
			goto tr19
		case 24:
			goto tr34
		case 25:
			goto tr36
		case 26:
			goto tr37
		case 27:
			goto tr37
		case 28:
			goto tr37
		case 29:
			goto tr37
		case 30:
			goto tr37
		case 31:
			goto tr37
		case 32:
			goto tr37
		case 33:
			goto tr37
		case 34:
			goto tr37
		case 35:
			goto tr37
		case 36:
			goto tr37
		case 37:
			goto tr37
		case 38:
			goto tr37
		case 39:
			goto tr37
		case 40:
			goto tr37
		case 41:
			goto tr37
		case 42:
			goto tr37
		case 43:
			goto tr37
		case 44:
			goto tr37
		}
	}

	_out: {}
	}

//line graphql_collections.rl:51
}

