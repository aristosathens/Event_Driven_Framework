# Aristos Athens

# Framework

# import threading as th
import multiprocessing as mp
import msvcrt
from Services import *





class Framework:
    def __init__(self, program_services, graphics_services):
        self.program_services = program_services
        self.graphics_services = graphics_services

        self.manager = mp.Manager()
        self.program_queue = self.manager.Queue()
        self.graphics_queue = self.manager.Queue()
        self.input_monitor = ProcessMonitor(self.program_queue, self.graphics_queue)
        self.program_event_handler = ProgramEventHandler(self.program_queue, self.program_services)
        self.graphics_event_handler = GraphicsEventHandler(self.graphics_queue, self.graphics_services)

        self.program_process = mp.Process(target=self.program_event_handler.run)
        self.graphics_process = mp.Process(target=self.graphics_event_handler.run)
        self.graphics_process.start()
        self.program_process.start()


    def run(self):

        new_event = Event(Event_Type.GLOBAL_START, None)
        self.program_queue.put(new_event)
        self.graphics_queue.put(new_event)
        self.input_monitor.run()

        self.program_process.join()
        self.graphics_process.join()


class ProcessMonitor:
    def __init__(self, program_queue, graphics_queue):
        self.program_queue = program_queue
        self.graphics_queue = graphics_queue


    def run(self):
        while(True):
            # print("here")
            if msvcrt.kbhit():
                new_event = Event(Event_Type.KEY_STROKE, msvcrt.getwch())
                self.program_queue.put(new_event)
                self.graphics_queue.put(new_event)



class ProgramEventHandler:
    def __init__(self, event_queue, services):
        # Create service objects from list of classes
        self.services = [service(event_queue) for service in services]
        self.event_queue = event_queue


    def run(self):
        while(True):
            # print("there")
            # if not self.event_queue.empty():
            event = self.event_queue.get()
            if event is None or event.event_type == Event_Type.NONE:
                continue
            elif event.event_type == Event_Type.GLOBAL_EXIT:
                break
            for service in self.services:
                if event.event_type in service.events_serviced:
                    return_event = service.run(event)
                    self.event_queue.put(return_event)

class GraphicsEventHandler:
    def __init__(self, event_queue, services):
        # Create service objects from list of classes
        self.services = [service(event_queue) for service in services]
        self.event_queue = event_queue



    def run(self):        
        while(True):
            # print("everywhere")
            # if not self.event_queue.empty():
            event = self.event_queue.get()
            if event is None or event.event_type == Event_Type.NONE:
                continue
            elif event.event_type == Event_Type.GLOBAL_EXIT:
                break
            for service in self.services:
                if event.event_type in service.events_serviced:
                    return_event = service.run(event)
                    self.event_queue.put(return_event)
                    continue









