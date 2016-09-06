module Players.Messages exposing (..)

import Http
import Players.Models exposing (PlayerID, Player)


type Msg
    = FetchAllDone (List Player)
    | FetchAllFail Http.Error
