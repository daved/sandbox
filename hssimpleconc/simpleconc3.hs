#!/usr/bin/env stack
-- stack --resolver lts-6.23 --install-ghc runghc --package async

import Control.Concurrent
import Control.Concurrent.Async

main:: IO()
main = do
   a <- async (do threadDelay 5000; putStrLn "Calculated result!"; return 42)
   r <- wait a
   putStrLn "Waiting..."
   putStrLn ("The answer is: " ++ show r)
