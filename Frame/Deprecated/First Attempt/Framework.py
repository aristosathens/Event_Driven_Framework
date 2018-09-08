# Aristos Athens

# Framework

# import threading as th
# import multiprocessing as mp
import msvcrt
from Services import *





class Framework:
    def __init__(self, program_services, graphics_services):
        self.program_services = program_services
        self.graphics_services = graphics_services

        self.manager = mp.Manager()
        self.program_queue = self.manager.Queue()
        self.program_event_handler = ProgramEventHandler(self.program_queue, self.program_services)
        self.graphics_event_handler = GraphicsEventHandler(self.program_queue, self.graphics_services)

        self.graphics_process = mp.Process(target=self.graphics_event_handler.run)



    def run(self):


        new_event = Event(Event_Type.GLOBAL_START, None)
        self.program_queue.put(new_event)

        self.graphics_process.start()
        self.program_event_handler.run()

        self.graphics_process.join()
        sys.exit(0)

class ProgramEventHandler:
    def __init__(self, event_queue, services):
        # Create service objects from list of classes
        self.services = [service(event_queue) for service in services]
        self.event_queue = event_queue


    def run(self):
        while(True):
            event = self.event_queue.get()
            if event is None or event.event_type == Event_Type.NONE:
                continue

            run = True
            if "program" in event.seen:
                run = False
            event.seen.append("program")
            if "graphics" not in event.seen:
                self.event_queue.put(event)
            if run == False:
                continue

            for service in self.services:
                if event.event_type in service.events_serviced:
                    return_event = service.run(event)
                    self.event_queue.put(return_event)

            if event.event_type == Event_Type.GLOBAL_EXIT:
                print("Exiting ProgramEventHandler")
                return


class GraphicsEventHandler:
    def __init__(self, event_queue, services):
        # Create service objects from list of classes
        self.services = [service(event_queue) for service in services]
        self.event_queue = event_queue

    def run(self):        
        while(True):
            event = self.event_queue.get()

            if event is None or event.event_type == Event_Type.NONE:
                continue

            run = True
            if "graphics" in event.seen:
                run = False
            event.seen.append("graphics")
            if "program" not in event.seen:
                self.event_queue.put(event)
            if run == False:
                continue

            for service in self.services:
                if event.event_type in service.events_serviced:
                    return_event = service.run(event)
                    self.event_queue.put(return_event)

            if event.event_type == Event_Type.GLOBAL_EXIT:
                print("Exiting GraphicsEventHandler")
                return









