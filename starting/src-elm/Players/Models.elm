module Players.Models exposing (..)


type alias PlayerID =
    Int


type alias Player =
    { id : PlayerID
    , name : String
    , level : Int
    }


new : Player
new =
    { id = 0
    , name = ""
    , level = 1
    }
