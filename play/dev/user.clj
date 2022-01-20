(ns user
  (:require [next.jdbc :as jdbc]
            [taoensso.timbre
             :refer [log  trace  debug  info  warn  error  fatal  report
                     logf tracef debugf infof warnf errorf fatalf reportf
                     spy get-env]]
            [sht.runner :as runner]
            [sht.build :as build]
            [sht.dkg-test :as dkg-test]
            [sht.play :as play]
            [sht.core :as core]))

(defn doit []
  (core/report-result [(runner/run-test (first @dkg-test/tests))]))

(comment
  (def db {:dbtype "postgresql"
           :dbname (play/keyper-db 0)
           :password core/play-db-password})

  (def ds (jdbc/get-datasource db))



  (jdbc/execute! ds ["select * from tendermint_batch_config"])
  (jdbc/execute-one! ds ["select * from meta_inf"])
  (jdbc/execute! ds ["select * from eons"])
  )
