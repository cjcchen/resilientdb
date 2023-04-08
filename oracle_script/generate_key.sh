#!/bin/sh

for idx in `seq 65 65`;
do
  echo `../bazel-bin/tools/key_generator_tools "./cert/node_${idx}" "ED25519"`
done
