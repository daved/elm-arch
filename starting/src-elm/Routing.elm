module Routing exposing (..)

import String
import Navigation
import UrlParser exposing (..)
import Players.Models exposing (PlayerID)


type Route
    = PlayersRoute
    | PlayerRoute PlayerID
    | NotFoundRoute


matchers : Parser (Route -> a) a
matchers =
    oneOf
        [ format PlayersRoute (s "")
        , format PlayerRoute (s "players" </> int)
        , format PlayersRoute (s "players")
        ]


hashParser : Navigation.Location -> Result String Route
hashParser location =
    location.hash
        |> String.dropLeft 1
        |> parse identity matchers


parser : Navigation.Parser (Result String Route)
parser =
    Navigation.makeParser hashParser


routeFromResult : Result String Route -> Route
routeFromResult result =
    case result of
        Ok route ->
            route

        Err string ->
            NotFoundRoute
