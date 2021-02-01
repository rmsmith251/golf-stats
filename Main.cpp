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
	//TODO
}

char mainMenu()
{
	char choice;

	cout << "Welcome to Ryan's Handicap Calculator (version 0.1)";
	cout << "\nPlease select an option from the list below.";
	cout << "\n1. Enter new round";
	cout << "\n2. Check handicap";
	cout << "\n3. View round stats";
	cout << "\n\n";
	cout << "\nPlease enter the number here -> ";
	fflush(stdin);
	choice = _getche();
	return choice;
}

int main() 
{
	while (1)
	{
		char selection;

		selection = mainMenu();

		switch (selection)
		{
		case '1':
			system("cls");
			cout << "Enter a new round";
			break;
		case '2':
			system("cls");
			cout << "Check Handicap";
			break;
		case '3':
			system("cls");
			cout << "View round stats";
			break;
		}
	}


	return 0;
}