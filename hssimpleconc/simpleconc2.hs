module Main where

import Control.Concurrent (forkIO, threadDelay)
import Control.Concurrent.MVar (newEmptyMVar, takeMVar, putMVar)

main :: IO ()
main = do
    result <- newEmptyMVar

    _ <- forkIO (do
        sleep 5
        putStrLn "Calculated result!"
        putMVar result 42)

    putStrLn "Waiting..."
    value <- takeMVar result
    putStrLn ("The answer is: " ++ show value)

sleep :: Int -> IO ()
sleep n = threadDelay (n * 1000)
