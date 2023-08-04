import trajectory_pb2
import spring_pb2
import matplotlib.pyplot as plt

trajectories = trajectory_pb2.Trajectories()

with open("../bin/trajectories_001001.bin", "rb") as file:
    trajectories.ParseFromString(file.read())

last_state_points = dict()

for trajectory in trajectories.trajectories:
    x, y = [], []

    for point in trajectory.points:
        x.append(point.PosX)
        y.append(point.PosY)

    last_state_points[trajectory.mass_id] = (x[-1], y[-1])

springs = spring_pb2.Springs()

with open("../bin/springs.bin", "rb") as file:
    springs.ParseFromString(file.read())

for spring in springs.springs:
    x1, y1 = last_state_points[spring.m1]
    x2, y2 = last_state_points[spring.m2]

    # plt.plot([x1, x2], [y1, y2])

for index in last_state_points.keys():
    x, y = last_state_points[index]
    plt.scatter(x, y)

plt.show()
