#include<iostream>
using namespace std;

//base class
class vehicle{
    public:
    string brand = "Ford";
    void honk(){
        cout << "Tuut,tuut \n" ; 
    }
    
};

//derived class
class Car : public Vehicle{
    public:
    string model = "Mustang";
};
int main(){
    Car myCar;
    myCar.honk();
    cout<< myCar.brand + " " + myCar.model;
    return 0;
}