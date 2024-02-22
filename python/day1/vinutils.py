"""
A bunch of utility functions created by Vinod

version 1.0
"""

def to_float(str_num):
    """
    converts the given str to a float, if possible and returns the same, else returns 0.0
    """
    try:
        return float(str_num)
    except ValueError:
        return 0.0


def dirr(obj):
    """
    this method returns non-dunder attributes of the given object
    """
    # attrs = []
    # for attr in dir(obj):
    #     if not attr.startswith('_'):
    #         attrs.append(attr)
    # return attrs
    return [each_attr for each_attr in dir(obj) if not each_attr.startswith('_')]

