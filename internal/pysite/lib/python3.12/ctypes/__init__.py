import struct

class c_float():
    def __init__(self, value):
        self.value = struct.unpack('f', struct.pack('f', value))
