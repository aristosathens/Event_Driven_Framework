# Aristos Athens

import time
import sys
import multiprocessing as mp
from PyQt5.QtCore import Qt as qt
from PyQt5.QtCore import QObject, pyqtSignal, pyqtSlot
from PyQt5.QtWidgets import QWidget, QDesktopWidget, QApplication

from Services_Template import *


# ------------------------------Program Services------------------------------ #


class Character_Service(ProgramService):
    def __init__(self, event_queue):
        events_serviced =   [
                            Event_Type.GLOBAL_START,
                            Event_Type.GLOBAL_EXIT,
                            Event_Type.KEY_DOWN,
                            Event_Type.KEY_UP
                            ]
        ProgramService.__init__(self, events_serviced, event_queue)
        self.position = (0, 0)

    def run(self, event):
        if event.parameter == 'w':
            self.position[0] -= 1
        if event.parameter == 's':
            self.position[0] += 1
        if event.parameter == 'a':
            self.position[1] -= 1
        if event.parameter == 'd':
            self.position[0] += 1



# ------------------------------Graphics Services------------------------------ #

# Service class. Spawns process for window (GUI) and does communication/compution in the original process
class Qt_Service(GraphicsService):
    def __init__(self, event_queue):
        events_serviced =   [
                            Event_Type.GLOBAL_START,
                            Event_Type.GLOBAL_EXIT,
                            Event_Type.UPDATE_GRAPHICS
                            ]
        GraphicsService.__init__(self, events_serviced, event_queue)

    def run(self, event):
        print("here")
        print(event.event_type)

        if event.event_type == Event_Type.GLOBAL_START:
            k = Qt_Wrapper(self.event_queue)
            self.gui_process = mp.Process(target=k.run, args=[event])
            self.gui_process.start()
            self.emitter = Qt_Emitter()

        elif event.event_type == Event_Type.UPDATE_GRAPHICS:
            print("Attempting to emit signal")
            self.emitter.emit(event)

        elif event.event_type == Event_Type.KEY_DOWN:
            key = event.parameter

class Qt_Emitter(QObject):
    my_signal = pyqtSignal(Event)
    def __init__(self):
        QObject.__init__(self, parent=None)
        self.my_signal.connect(Qt_Window.respond_to_event)

    def emit(self, event):
        self.my_signal.emit(event)



# Wrapper class that creates and runs the window object.
# Main reason for this is that we can't pickle a QObject, so we can't set a Process target to app.exe_() directly 
class Qt_Wrapper:
    def __init__(self, event_queue):
        self.event_queue = event_queue

    def run(self, event):
        if event.event_type == Event_Type.GLOBAL_START:
            self.app = QApplication(sys.argv)
            self.window = Qt_Window(self.event_queue)
            self.app.exec_()

# Window class. Actual GUI processes occur here
class Qt_Window(QWidget):
    def __init__(self, event_queue):
        super().__init__()
        self.width = 400
        self.height = 300
        self.resize(self.width, self.height)
        self.center_window()
        self.event_queue = event_queue
        self.setWindowTitle("Aristos' App")
        self.show()


    def center_window(self):
        self.desktop_center = QDesktopWidget().availableGeometry().center()
        self.desktop_center.setX(self.desktop_center.x() - int(self.width/2))
        self.desktop_center.setY(self.desktop_center.y() - int(self.height/2))
        self.move(self.desktop_center)

    @pyqtSlot(Event)
    def respond_to_event(self, event):
        self.move(350, 700)
        print("Signal received")
        # print("Qt_Window received external event: ", event)

    def keyPressEvent(self, e):
        if e.isAutoRepeat():
            return

        key = e.key()
        if key == qt.Key_Escape:
            self.event_queue.put(Event(Event_Type.GLOBAL_EXIT, None))
            self.close()
            sys.exit(0)
            return
        elif key == qt.Key_C:
            self.move(200, 600)
        elif key == qt.Key_F:
            self.move(600, 200)

        elif key == qt.Key_E:
            self.event_queue.put(Event(Event_Type.UPDATE_GRAPHICS, None))

        elif key < 0x10FFFF:
            self.event_queue.put(Event(Event_Type.KEY_DOWN, chr(key)))
