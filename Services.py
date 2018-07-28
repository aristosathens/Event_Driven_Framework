# Aristos Athens

import enum
import time
import pygame as pg



Event_Type = enum.Enum("Event_Type", [ "NONE", "GLOBAL_START", "GLOBAL_EXIT", "KEY_DOWN", "KEY_UP", "POLL_PYGAME" ])

# Event_Typeclass. Messages are passed around using these
class Event:
    def __init__(self, event_type, parameter):
        self.event_type = event_type
        self.parameter = parameter

# All service objects require init() and run() functions
# A service is initialized with the list of Events it responds to
# The run() function takes an Event_Typeobject and responds accordingly

# Super class for all services dealing with internal app logic
class ProgramService:
    def __init__(self, events_serviced, event_queue):
        self.event_queue = event_queue
        self.events_serviced = events_serviced

class Keystroke_Service(ProgramService):
    def __init__(self, event_queue):
        events_serviced =   [
                            Event_Type.KEY_DOWN,
                            Event_Type.KEY_UP
                            ]
        ProgramService.__init__(self, events_serviced, event_queue)
        print("events serviced: ", self.events_serviced)

    def run(self, event):
        print("Keystroke: ", event.parameter)





# Super class for all services dealing with graphics and windowing
class GraphicsService:
    def __init__(self, events_serviced, event_queue):
        self.event_queue = event_queue
        self.events_serviced = events_serviced
        pass

class PyGame_Service(GraphicsService):
    def __init__(self, event_queue):
        events_serviced =   [
                            Event_Type.GLOBAL_START,
                            Event_Type.POLL_PYGAME,
                            Event_Type.KEY_DOWN,
                            Event_Type.KEY_UP
                            ]
        GraphicsService.__init__(self, events_serviced, event_queue)
        print("events serviced: ", self.events_serviced)

    def run(self, event):

        if event.event_type == Event_Type.GLOBAL_START:
            print("here")
            self.window_width = 400
            self.window_height = 300
            pg.init()
            pg.display.set_mode((self.window_width, self.window_height))
            return(Event(Event_Type.POLL_PYGAME,None))


        elif event.event_type == Event_Type.POLL_PYGAME:
            # print(time.time())
            # print("POLLING")
            for next_event in pg.event.get():
                print(next_event)
                if next_event.type == pg.KEYDOWN:
                    print("KEY DOWN")
                    self.event_queue.put(Event(Event_Type.KEY_DOWN, next_event.key))
                elif next_event.type == pg.KEYUP:
                    print("KEY UP")
                    self.event_queue.put(Event(Event_Type.KEY_UP, next_event.key))


            return(Event(Event_Type.POLL_PYGAME,None))

