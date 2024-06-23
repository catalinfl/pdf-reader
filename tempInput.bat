@echo off 
color 07 
title pdfreader
:loop
set /p UserInput=
if "%UserInput%"=="exit" goto :end
echo You entered: %UserInput%
goto :loop
:end
pause
