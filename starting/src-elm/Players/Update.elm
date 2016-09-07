module Players.Update exposing (..)

import Players.Commands exposing (save)
import Players.Messages exposing (Msg(..))
import Players.Models exposing (Player, PlayerID)
import Navigation


changeLevelCommands : PlayerID -> Int -> List Player -> List (Cmd Msg)
changeLevelCommands playerID howMuch =
    let
        cmdForPlayer existingPlayer =
            if existingPlayer.id == playerID then
                save { existingPlayer | level = existingPlayer.level + howMuch }
            else
                Cmd.none
    in
        List.map cmdForPlayer


updatePlayer : Player -> List Player -> List Player
updatePlayer updatedPlayer =
    let
        select existingPlayer =
            if existingPlayer.id == updatedPlayer.id then
                updatedPlayer
            else
                existingPlayer
    in
        List.map select


update : Msg -> List Player -> ( List Player, Cmd Msg )
update message players =
    case message of
        FetchAllDone newPlayers ->
            ( newPlayers, Cmd.none )

        FetchAllFail error ->
            ( players, Cmd.none )

        ShowPlayers ->
            ( players, Navigation.newUrl "#players" )

        ShowPlayer id ->
            ( players, Navigation.newUrl ("#players/" ++ (toString id)) )

        ChangeLevel id howMuch ->
            ( players, changeLevelCommands id howMuch players |> Cmd.batch )

        SaveSuccess updatedPlayer ->
            ( updatePlayer updatedPlayer players, Cmd.none )

        SaveFail error ->
            ( players, Cmd.none )
