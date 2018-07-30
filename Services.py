# Aristos Athens

import tkinter as tk
from tkinter import ttk

from Services_Template import *


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




class Tk_Service(GraphicsService):
    def __init__(self, event_queue):
        events_serviced =   [
                            Event_Type.GLOBAL_START,
                            Event_Type.GLOBAL_EXIT,
                            Event_Type.UPDATE_GRAPHICS
                            ]
        GraphicsService.__init__(self, events_serviced, event_queue)

    def run(self, event):

        if event.event_type == Event_Type.GLOBAL_START:
            root = tk.Tk()
            root.title("Feet to Meters")
            mainframe = ttk.Frame(root, padding="3 3 12 12")
            mainframe.grid(column=0, row=0, sticky="N, W, E, S")
            mainframe.columnconfigure(0, weight=1)
            mainframe.rowconfigure(0, weight=1)

            self.feet = tk.StringVar()
            self.meters = tk.StringVar()
            feet_entry = ttk.Entry(mainframe, width=7, textvariable=self.feet)
            feet_entry.grid(column=2, row=1, sticky="W, E")
            ttk.Label(mainframe, textvariable=self.meters).grid(column=2, row=2, sticky="W, E")
            ttk.Button(mainframe, text="Calculate", command=self.calculate).grid(column=3, row=3, sticky="W")
            ttk.Label(mainframe, text="feet").grid(column=3, row=1, sticky="W")
            ttk.Label(mainframe, text="is equivalent to").grid(column=1, row=2, sticky="E")
            ttk.Label(mainframe, text="meters").grid(column=3, row=2, sticky="W")

            for child in mainframe.winfo_children():
                child.grid_configure(padx=5, pady=5)
            feet_entry.focus()
            root.bind('<Return>', self.calculate)
            root.mainloop()

    def calculate(self, *args):
        try:
            value = float(feet.get())
            self.meters.set((0.3048 * value * 10000.0 + 0.5)/10000.0)
        except ValueError:
            pass
            # elif event.event_type == Event_Type.UPDATE_GRAPHICS:


        # elif event.event_type == Event_Type.KEY_DOWN:
        #     key = event.parameter