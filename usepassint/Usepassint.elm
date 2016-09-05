module Usepassint exposing (..)

import Html exposing (Html)
import Html.App
import Widgetpassint


-- MODEL


type alias AppModel =
    { widgetModel : Widgetpassint.Model
    }


initModel : AppModel
initModel =
    { widgetModel = Widgetpassint.initModel
    }


init : ( AppModel, Cmd Msg )
init =
    ( initModel, Cmd.none )



-- MESSAGES


type Msg
    = WidgetMsg Widgetpassint.Msg



-- VIEW


view : AppModel -> Html Msg
view model =
    Html.div []
        [ Html.App.map WidgetMsg (Widgetpassint.view model.widgetModel)
        ]



-- UPDATE


update : Msg -> AppModel -> ( AppModel, Cmd Msg )
update message model =
    case message of
        WidgetMsg subMsg ->
            let
                ( updatedWidgetModel, widgetCmd ) =
                    Widgetpassint.update subMsg model.widgetModel
            in
                ( { model | widgetModel = updatedWidgetModel }, Cmd.map WidgetMsg widgetCmd )



-- SUBSCRIPTIONS


subscriptions : AppModel -> Sub Msg
subscriptions model =
    Sub.none



-- APP


main : Program Never
main =
    Html.App.program
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        }
