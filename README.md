# attendance-api
golang attendance api


TO RUN
go run main.go

TO TEST 
go test ./...

ROADMAP
1. *show a table full list of people*

in react - show a list (table at first)


api:
GetMembers:
return 100 (paginate? maybe...but not highest priority)
DONE

GetAttendance
given a member id return that members attendance record
GetAttendance(member_id)
return []attendance list of events they attended
DONE

GetMembersByDay
if valid date, return list of members
if not, return error for that day

AttendanceCount 
returns a date and the member count associated
