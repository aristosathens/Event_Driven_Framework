# Aristos Athens

import enum


Event_Type = enum.Enum("Event_Type",    [
                                        "NONE",
                                        "GLOBAL_START",
                                        "GLOBAL_EXIT",
                                        "KEY_DOWN",
                                        "KEY_UP",
                                        "UPDATE_GRAPHICS"
                                        ])

# Event_Typeclass. Messages are passed around using these
class Event:
    def __init__(self, event_type, parameter):
        self.event_type = event_type
        self.parameter = parameter
        self.seen = []

# All service objects require init() and run() functions
# A service is initialized with the list of Events it responds to
# The run() function takes an Event_Typeobject and responds accordingly

# Super class for all services dealing with internal app logic
class ProgramService:
    def __init__(self, events_serviced, event_queue):
        self.event_queue = event_queue
        self.events_serviced = events_serviced

    def run(self):
        pass


# Super class for all services dealing with graphics and windowing
class GraphicsService:
    def __init__(self, events_serviced, event_queue):
        self.event_queue = event_queue
        self.events_serviced = events_serviced
        pass

    def run(self):
        pass