#include <iostream>
#include <cstdio>
#include <cstring>
#include <cstdlib>
#include <conio.h>
#include <iomanip>
#include "Source.h"

using namespace std;

void roundStats()
{
	//TODO
}

void checkHandicap() 
{
	//TODO
}

void newRound()
{


	char choice;

	cout << ("\n\nEnter y to go to the main menu or press any key to exit."); //TODO: add submenu to go to other functions
	choice = _getche();

	if (choice == 'y' || choice == 'Y')
	{
		mainMenu();
	}
}

void mainMenu()
{
	char choice;
	system("cls");
	cout << "Welcome to Ryan's Handicap Calculator (version 0.1)";
	cout << "\nPlease select an option from the list below.";
	cout << "\n1. Enter new round";
	cout << "\n2. Check handicap";
	cout << "\n3. View round stats";
	cout << "\n4. Exit";
	cout << "\n\n";
	cout << "\nPlease enter the number here -> ";
	fflush(stdin);
	choice = _getche();

	switch (choice)
	{
	case '1':
		system("cls");
		cout << "Enter a new round";
		newRound();
		break;
	case '2':
		system("cls");
		cout << "Check Handicap";
		checkHandicap();
		break;
	case '3':
		system("cls");
		cout << "View round stats";
		roundStats();
		break;
	case '4':
		system("cls");
		cout << "Goodbye!";
		break;
	}
}

int main() 
{
	mainMenu();

	return 0;
}