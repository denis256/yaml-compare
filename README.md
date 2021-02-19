# Yaml files compare

Minimalistic go application to compare if keys from "custom-values.yaml"
override keys "values.yaml".

Application will print in the end list of not matched keys by "custom-values.yaml"

Example output based on sample files:
```
***Values FlatMap:
.key3.config3 => 666
.key3.config4 => 111
.key1.key2.config1 => value1
.key1.key2.qwe => abc
.key1.config2 => value2
.key1.xyz => 123

***Custom FlatMap:
.key1.xyz1 => 123
.key3.config666 => 111
.key3.config4 => 111
.key1.key2.config1 => value1
.key1.key2.custom1.value2 => qwe

***Keys from custom yaml which not overriding values.yaml:
.key1.key2.custom1.value2
.key1.xyz1
.key3.config666

```