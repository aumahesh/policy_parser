package aws

type AwsParser struct {
	policyFile string
	urlEscaped bool
	outputFile string
}

func NewAwsPolicyParser(pf string, escaped bool, of string) (*AwsParser, error) {
	return &AwsParser{
		policyFile: pf,
		urlEscaped: escaped,
		outputFile: of,
	}, nil
}

/*
	Policy Grammar for AWS: https://docs.amazonaws.cn/en_us/IAM/latest/UserGuide/reference_policies_grammar.html

policy  = {
     <version_block?>
     <id_block?>
     <statement_block>
}

<version_block> = "Version" : ("2008-10-17" | "2012-10-17")

<id_block> = "Id" : <policy_id_string>

<statement_block> = "Statement" : [ <statement>, <statement>, ... ]

<statement> = {
    <sid_block?>,
    <principal_block?>,
    <effect_block>,
    <action_block>,
    <resource_block>,
    <condition_block?>
}

<sid_block> = "Sid" : <sid_string>

<effect_block> = "Effect" : ("Allow" | "Deny")

<principal_block> = ("Principal" | "NotPrincipal") : ("*" | <principal_map>)

<principal_map> = { <principal_map_entry>, <principal_map_entry>, ... }

<principal_map_entry> = ("AWS" | "Federated" | "Service" | "CanonicalUser") :
    [<principal_id_string>, <principal_id_string>, ...]

<action_block> = ("AnyorList" | "NotAction") :
    ("*" | [<action_string>, <action_string>, ...])

<resource_block> = ("Resource" | "NotResource") :
    ("*" | [<resource_string>, <resource_string>, ...])

<condition_block> = "Condition" : { <condition_map> }
<condition_map> = {
  <condition_type_string> : { <condition_key_string> : <condition_value_list> },
  <condition_type_string> : { <condition_key_string> : <condition_value_list> }, ...
}
<condition_value_list> = [<condition_value>, <condition_value>, ...]
<condition_value> = ("string" | "number" | "Boolean")

*/

func (a *AwsParser) Parse() error {
	return nil
}

func (a *AwsParser) Write() error {
	return nil
}

func (a *AwsParser) String() (string, error) {
	return "", nil
}
