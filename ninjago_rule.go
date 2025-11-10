package ninjago

func MakeNinjaGoCXXRule(tool string) *NinjaGoRuleBase {
	return MakeNinjaGoRule(
		"cxx",
		(tool + " -c $in -o $out $cflags $includes"),
		"CXX $out",
	)
}

func MakeNinjaGoCCrule(tool string) *NinjaGoRuleBase {
	return MakeNinjaGoRule(
		"cc",
		(tool + " -c $in -o $out $cflags $includes"),
		"CC $out",
	)
}

func MakeNinjaGoLDRule(tool string) *NinjaGoRuleBase {
	return MakeNinjaGoRule(
		"link",
		(tool + " $in -o $out $ldflags"),
		"LINK $out",
	)
}

func MakeNinjaGoCleanRule() *NinjaGoRuleBase {
	return MakeNinjaGoRule(
		"clean",
		"rm -rf $in",
		"CLEAN $in",
	)
}
