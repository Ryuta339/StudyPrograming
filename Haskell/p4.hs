-- problem 31
isPrime :: Integeral a => a -> Bool
isPrime n = isPrime' [1..n] n
