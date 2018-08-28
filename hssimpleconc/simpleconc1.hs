import System.IO
import Control.Concurrent
import Control.Concurrent.Chan
import Text.Printf
 
type Pipe = Chan (Either String String)
 
main :: IO ()
main = do
    chan      <- newChan :: IO Pipe
    s         <- getChanContents chan   -- lazy list of chan elements
    c1Thread  <- forkIO $ reader "c1" (catLeft  s) -- read only Lefts
    c2Thread  <- forkIO $ reader "c2" (catRight s) -- read only Rights
    writer chan
  where
    catLeft  ls = [x | Left  x <- ls]
    catRight ls = [x | Right x <- ls]
 
writer :: Pipe -> IO ()
writer chan = loop
  where
    loop = getChar >>= command
    command '0'  = print "done"
    command '1'  = writeChan chan (Left  "main: 1") >> loop
    command '2'  = writeChan chan (Right "main: 2") >> loop
    command '\n' = loop -- ignore
    command c    = printf "Illegal: %c\n" c         >> loop
 
reader :: String -> [String] -> IO ()
reader name xs = mapM_ (printf "%s %s\n" name) xs
