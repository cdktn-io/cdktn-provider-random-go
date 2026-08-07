// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package ephemeralrandompassword

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-provider-random-go/random/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktn-io/cdktn-provider-random-go/random/v15/ephemeralrandompassword/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/random/3.9.0/docs/ephemeral-resources/password random_password}.
type EphemeralRandomPassword interface {
	cdktn.TerraformEphemeralResource
	BcryptHash() *string
	// Experimental.
	CdktfStack() cdktn.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktn.ITerraformIterator
	// Experimental.
	SetForEach(val cdktn.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Length() *float64
	SetLength(val *float64)
	LengthInput() *float64
	// Experimental.
	Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle)
	Lower() interface{}
	SetLower(val interface{})
	LowerInput() interface{}
	MinLower() *float64
	SetMinLower(val *float64)
	MinLowerInput() *float64
	MinNumeric() *float64
	SetMinNumeric(val *float64)
	MinNumericInput() *float64
	MinSpecial() *float64
	SetMinSpecial(val *float64)
	MinSpecialInput() *float64
	MinUpper() *float64
	SetMinUpper(val *float64)
	MinUpperInput() *float64
	// The tree node.
	Node() constructs.Node
	Numeric() interface{}
	SetNumeric(val interface{})
	NumericInput() interface{}
	OverrideSpecial() *string
	SetOverrideSpecial(val *string)
	OverrideSpecialInput() *string
	// Experimental.
	Provider() cdktn.TerraformProvider
	// Experimental.
	SetProvider(val cdktn.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Result() *string
	Special() interface{}
	SetSpecial(val interface{})
	SpecialInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Upper() interface{}
	SetUpper(val interface{})
	UpperInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Registers a synth-time validation that the project's declared targetVersions admit the given provider-protocol feature family.
	//
	// Called by generated provider bindings when a versioned feature is
	// structurally in use - the element's existence in the construct tree
	// already implies the feature is used, e.g. constructing a
	// `TerraformEphemeralResource` at all - so, unlike
	// `_registerResolveDiscoveredProviderFeatureUsage`, this registration is
	// never deactivated by `_resetResolveDiscoveredProviderFeatureUsage`. Not
	// intended to be called directly by user code. Lives on `TerraformElement`
	// (rather than `TerraformResource`) so it covers any element subclass
	// that needs it.
	// Experimental.
	RegisterProviderFeatureUsage(feature cdktn.ProviderFeature)
	ResetLower()
	ResetMinLower()
	ResetMinNumeric()
	ResetMinSpecial()
	ResetMinUpper()
	ResetNumeric()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrideSpecial()
	ResetSpecial()
	ResetUpper()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this ephemeral resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
	// Applies one or more mixins to this construct.
	//
	// Mixins are applied in order. The list of constructs is captured at the
	// start of the call, so constructs added by a mixin will not be visited.
	// Use multiple `with()` calls if subsequent mixins should apply to added
	// constructs.
	//
	// Returns: This construct for chaining.
	With(mixins ...constructs.IMixin) constructs.IConstruct
}

// The jsii proxy struct for EphemeralRandomPassword
type jsiiProxy_EphemeralRandomPassword struct {
	internal.Type__cdktnTerraformEphemeralResource
}

func (j *jsiiProxy_EphemeralRandomPassword) BcryptHash() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bcryptHash",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) CdktfStack() cdktn.TerraformStack {
	var returns cdktn.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) ForEach() cdktn.ITerraformIterator {
	var returns cdktn.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) LengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Lifecycle() *cdktn.TerraformEphemeralResourceLifecycle {
	var returns *cdktn.TerraformEphemeralResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Lower() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lower",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) LowerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lowerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinLower() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLower",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinLowerInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minLowerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinNumeric() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumeric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinNumericInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minNumericInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinSpecial() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSpecial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinSpecialInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minSpecialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinUpper() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) MinUpperInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minUpperInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Numeric() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numeric",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) NumericInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numericInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) OverrideSpecial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideSpecial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) OverrideSpecialInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"overrideSpecialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Provider() cdktn.TerraformProvider {
	var returns cdktn.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Result() *string {
	var returns *string
	_jsii_.Get(
		j,
		"result",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Special() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"special",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) SpecialInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specialInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) TerraformGeneratorMetadata() *cdktn.TerraformProviderGeneratorMetadata {
	var returns *cdktn.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) Upper() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upper",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EphemeralRandomPassword) UpperInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"upperInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/random/3.9.0/docs/ephemeral-resources/password random_password} Ephemeral Resource.
func NewEphemeralRandomPassword(scope constructs.Construct, id *string, config *EphemeralRandomPasswordConfig) EphemeralRandomPassword {
	_init_.Initialize()

	if err := validateNewEphemeralRandomPasswordParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EphemeralRandomPassword{}

	_jsii_.Create(
		"@cdktn/provider-random.ephemeralRandomPassword.EphemeralRandomPassword",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/random/3.9.0/docs/ephemeral-resources/password random_password} Ephemeral Resource.
func NewEphemeralRandomPassword_Override(e EphemeralRandomPassword, scope constructs.Construct, id *string, config *EphemeralRandomPasswordConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-random.ephemeralRandomPassword.EphemeralRandomPassword",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetForEach(val cdktn.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetLength(val *float64) {
	if err := j.validateSetLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"length",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetLifecycle(val *cdktn.TerraformEphemeralResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetLower(val interface{}) {
	if err := j.validateSetLowerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lower",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetMinLower(val *float64) {
	if err := j.validateSetMinLowerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minLower",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetMinNumeric(val *float64) {
	if err := j.validateSetMinNumericParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minNumeric",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetMinSpecial(val *float64) {
	if err := j.validateSetMinSpecialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minSpecial",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetMinUpper(val *float64) {
	if err := j.validateSetMinUpperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minUpper",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetNumeric(val interface{}) {
	if err := j.validateSetNumericParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numeric",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetOverrideSpecial(val *string) {
	if err := j.validateSetOverrideSpecialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideSpecial",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetProvider(val cdktn.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetSpecial(val interface{}) {
	if err := j.validateSetSpecialParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"special",
		val,
	)
}

func (j *jsiiProxy_EphemeralRandomPassword)SetUpper(val interface{}) {
	if err := j.validateSetUpperParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upper",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func EphemeralRandomPassword_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralRandomPassword_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-random.ephemeralRandomPassword.EphemeralRandomPassword",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralRandomPassword_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralRandomPassword_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-random.ephemeralRandomPassword.EphemeralRandomPassword",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EphemeralRandomPassword_IsTerraformEphemeralResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEphemeralRandomPassword_IsTerraformEphemeralResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktn/provider-random.ephemeralRandomPassword.EphemeralRandomPassword",
		"isTerraformEphemeralResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EphemeralRandomPassword_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktn/provider-random.ephemeralRandomPassword.EphemeralRandomPassword",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) RegisterProviderFeatureUsage(feature cdktn.ProviderFeature) {
	if err := e.validateRegisterProviderFeatureUsageParameters(feature); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"registerProviderFeatureUsage",
		[]interface{}{feature},
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetLower() {
	_jsii_.InvokeVoid(
		e,
		"resetLower",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetMinLower() {
	_jsii_.InvokeVoid(
		e,
		"resetMinLower",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetMinNumeric() {
	_jsii_.InvokeVoid(
		e,
		"resetMinNumeric",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetMinSpecial() {
	_jsii_.InvokeVoid(
		e,
		"resetMinSpecial",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetMinUpper() {
	_jsii_.InvokeVoid(
		e,
		"resetMinUpper",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetNumeric() {
	_jsii_.InvokeVoid(
		e,
		"resetNumeric",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetOverrideSpecial() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideSpecial",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetSpecial() {
	_jsii_.InvokeVoid(
		e,
		"resetSpecial",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) ResetUpper() {
	_jsii_.InvokeVoid(
		e,
		"resetUpper",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EphemeralRandomPassword) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EphemeralRandomPassword) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		e,
		"with",
		args,
		&returns,
	)

	return returns
}

