# payment_user_svc_exercise

Code exercise: A service that creates and retrieves users from an internal store and two payment gateways.

## Installation

* Install Go 1.14 if needed
* Note that all dependencies are vendored, so nothing else should need installation.
* In the project root, run `STRIPE_KEY=... BRAINTREE_MERCHANT_ID=... BRAINTREE_PUBLIC_KEY=... BRAINTREE_PRIVATE_KEY=... make run` (replace each `...` with a real value)
* The API is now available at `http://127.0.0.1:8000`.
* In another shell, run `make test` to run the integration tests.
