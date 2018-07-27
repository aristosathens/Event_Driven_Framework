# Aristos Athens

# Framework

# import threading as th
import multiprocessing as mp
import pygame as pg
import enum
import msvcrt


Events = enum.Enum("Events", ["NONE", "KEY_STROKE"])

class Framework:
    def __init__(self):
        self.services = [
                        Keystroke_Service(),
                        Display_Service(),
                        ]
        self.event_queue = mp.Queue()
        self.input_monitor = InputMonitor(self.event_queue)
        self.event_handler = EventHandler(self.event_queue, self.services)



        self.handler_process = mp.Process(target=self.event_handler.run)
        self.handler_process.start()

        self.input_monitor.run()


class InputMonitor:
    def __init__(self, event_queue):
        self.event_queue = event_queue

    def run(self):
        while(True):
            # print("here")
            if msvcrt.kbhit():
                new_event = Event(Events.KEY_STROKE, msvcrt.getwch())
                self.event_queue.put(new_event)


class EventHandler:
    def __init__(self, event_queue, services):
        self.services = services
        self.event_queue = event_queue

    def run(self):
        while(True):
            # print("there")
            if not self.event_queue.empty():
                event = self.event_queue.get()
                if event.event_type == Events.NONE:
                    continue
                for service in services:
                    if event.event_type in service.events_serviced:
                        return_event = service.post(event)
                        event_queue.put(return_event)



class Event:
    def __init__(self, event_type, parameter):
        self.event_type = event_type
        self.parameter = parameter

class Service:
    def __init__(self, events_serviced):
        self.events_serviced = events_serviced
        pass

    def post(self, event):
        pass

class Keystroke_Service(Service):
    def __init__(self):
        events_serviced = [Events.KEY_STROKE]
        Service.__init__(self, events_serviced)

    def post(self, event):
        print("Keystroke: ", event.parameter)

class Display_Service(Service):
    def __init__(self):
        pass

    def post(self, event):
        pass

# class Character_Service(Service):




if __name__ == "__main__":
    Framework = Framework()