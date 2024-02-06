(use-modules
  ((gnu packages golang) #:select (go-1.21))
  ((guix build-system gnu) #:select (gnu-build-system))
  ((guix licenses) #:prefix licenses/)
  ((guix packages) #:select (package)))

(package
  (name "jupiter-subscribed")
  (version "git")
  (source #f)
  (build-system gnu-build-system)
  (native-inputs
    (list
      go-1.21))
  (synopsis "Jupiter Cloud: SubscribeD")
  (description "Jupiter Cloud: SubscribeD")
  (home-page "https://www.jupitercloud.com")
  (license licenses/asl2.0))
