@echo off 
color 61 
@echo off 
color 44 
@echo off 
color 66 
title pdfreader
:loop
set /p UserInput=
if "%UserInput%"=="exit" goto :end
echo You entered: %UserInput%
goto :loop
:end

pause
