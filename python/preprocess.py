import random

import mass_pb2
import spring_pb2


def read_file(file_path="../data/weighted_network_big.txt"):
    entries = []
    with open(file_path, "r") as file:
        for line in file.readlines():
            m1, m2, weight = line.replace("\n", "").split(" ")

            entries.append([int(m1), int(m2), float(weight)])

    return entries


def main():
    network = read_file()
    nodes = dict()
    edges = dict()

    masses = mass_pb2.Masses()
    springs = spring_pb2.Springs()

    for entry in network:
        nodes[entry[0]] = 0
        # TODO: Allow for other information about the spring,
        #  currently only one value for relaxed length is allowed
        edges[frozenset(entry[:2])] = entry[2]

    for index in nodes.keys():
        mass = mass_pb2.Mass()
        mass.id = index
        mass.pos_x = random.random()
        mass.pos_y = random.random()
        mass.weight = 1
        mass.drag = 0.95
        mass.fixed = False

        masses.masses.append(mass)

    for index, item in enumerate(edges):
        x, y = item

        spring = spring_pb2.Spring()
        spring.id = index
        spring.m1 = x
        spring.m2 = y
        spring.relaxed_length = edges[item] / 100000
        spring.spring_constant = 0.05

        springs.springs.append(spring)

    print(len(springs.springs))

    with open("../bin/masses.bin", "wb") as f:
        f.write(masses.SerializeToString())

    with open("../bin/springs.bin", "wb") as f:
        f.write(springs.SerializeToString())


main()
