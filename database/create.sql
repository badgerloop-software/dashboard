# DDL for the dashboard database
USE Dashboard;

# Drop tables upon create in case they exist
DROP TABLE IF EXISTS `Update`;
DROP TABLE IF EXISTS `Data`;

# Data table, holds all of the specified fields
CREATE TABLE `Data` (
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

	# Required SpaceX arguments
	`team_id` INTEGER(3) DEFAULT 55,		# UINT8 - Team Identifier (TODO)
	`status` INTEGER(3) NOT NULL,			# UINT8 - Pod Status *
	`acceleration` INTEGER(10) NOT NULL,	# INT32 - Acceleration in cm / s^2
	`position` INTEGER(10) NOT NULL,		# INT32 - Position in cm
	`velocity` INTEGER(10) NOT NULL,		# INT32 - Velocity in cm/s

	# Optional SpaceX arguments
	`battery_voltage` INTEGER(10) NOT NULL,		# INT32 - Battery voltage in mV
	`battery_current` INTEGER(10) NOT NULL,		# INT32 - Battery current in mA
	`battery_temperature` INTEGER(10) NOT NULL,	# INT32 - Battery temp. in tenths deg. C
	`pod_temperature` INTEGER(10) NOT NULL,		# INT32 - Pod temp. in tenths deg. C
	`stripe_count` INTEGER(10) NOT NULL,		# UINT32 - Count of optical navigation stripes

	# Additional fields for dashboard
	`pod_pressure` INTEGER(5) NOT NULL,			# UINT16 - pascals?
	`switch_states` INTEGER(5) NOT NULL,		# UINT16 - bit fields, see badgerloop.h
	`pr_p1` INTEGER(5) NOT NULL,				# UINT16 - PSI, TODO: which is this?
	`pr_p2` INTEGER(5) NOT NULL,				# UINT16 - PSI, TODO: which is this?
	`br_p1` INTEGER(3) NOT NULL,				# UINT16 - PSI, TODO: which is this?
	`br_p2` INTEGER(3) NOT NULL,				# UINT16 - PSI, TODO: which is this?
	`br_p3` INTEGER(3) NOT NULL					# UINT16 - PSI, TODO: which is this?
);

# Update table
# Keeps track of all of the updates in sequence
CREATE TABLE `Update` (
	`id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
	`created` TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
	`last_update_start` INTEGER(10) NOT NULL,
	`last_update_end` INTEGER(10) NOT NULL,
	FOREIGN KEY(`last_update_start`) REFERENCES `Data`(`id`),
	FOREIGN KEY(`last_update_end`) REFERENCES `Data`(`id`)
);

# * Pod Status:
# 0: Fault - If seen, will cause SpaceX to abort the tube run
# 1: Idle - Any state where the pod is on, but not ready to be pushed
# 2: Ready - Any state where the pod is ready to be pushed
# 3: Pushing - Any state where the pod detects it is being pushed
# 4: Coast - Any state when the pod detects it has separated from the pusher vehicle
# 5: Braking - Any state when the pod is applying its brakes

