(ns {{.Namespace}}.core-test
  (:require [clojure.test :refer :all]
            [{{.Namespace}}.core :as core]))

(deftest greet-test
  (is (= "Hello, World!" (core/greet "World"))))