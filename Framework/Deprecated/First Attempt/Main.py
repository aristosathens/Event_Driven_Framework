# Aristos Athens

import Framework
import Services


if __name__ == "__main__":
    
    program_services = Services.ProgramService.__subclasses__()
    graphics_services = Services.GraphicsService.__subclasses__()
    print("Program services: ", program_services)
    print("Graphics services: ", graphics_services)

    Framework = Framework.Framework(program_services, graphics_services)
    Framework.run()